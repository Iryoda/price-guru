package jobs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/imroc/req/v3"
	"github.com/iryoda/price-guru/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/html"
)

type Listener struct{}

type HtmlNodeInfo struct {
	ParentTag       string
	Content         string
	QueryableString string
}

func (l Listener) Listen() {
	ch := InitAmqp()
	c := InitMongo()
	db := c.Database("price-guru").Collection("watchers")

	msg, err := ch.Consume(
		os.Getenv("WATCHER_QUEUE"),
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	fmt.Println("Listening for messages in watchers QUEUE...")

	fakeChrome := req.DefaultClient().ImpersonateChrome()
	cc := colly.NewCollector(
		colly.UserAgent(fakeChrome.Headers.Get("User-Agent")),
	)
	cc.WithTransport(fakeChrome.Transport)

	go func() {
		for d := range msg {
			var watcher entities.Watcher
			err = db.FindOne(context.TODO(), bson.M{"_id": d.Body}).Decode(&watcher)

			if err != nil {
				log.Println(err)
				continue
			}

			price, err := CheckPrice(cc, watcher)

			if err != nil {
				log.Println(err)
				continue
			}

			update := bson.D{{Key: "$set", Value: bson.D{
				{Key: "status", Value: entities.SUCCESS},
				{Key: "lastRun", Value: time.Now()},
				{Key: "lastValue", Value: price},
			}}}

			db.UpdateOne(
				context.TODO(),
				bson.M{"_id": d.Body},
				update,
			)
		}
	}()
	<-forever
}

func getNode(n *html.Node) *html.Node {
	var result *html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data != "html" &&
			c.Data != "head" &&
			c.Data != "body" {
			return c
		}

		result = getNode(c)
	}

	return result
}

func CreateQueryableString(n *html.Node) string {
	tag := n.Data

	for _, a := range n.Attr {
		tag += fmt.Sprintf(`[%s='%s']`, a.Key, a.Val)
	}

	return strings.TrimSpace(tag)
}

func GetNodeQueryableHtmlWithParser(tag string) (HtmlNodeInfo, error) {
	n, err := html.Parse(strings.NewReader(tag))

	if err != nil {
		return HtmlNodeInfo{}, err
	}

	a := getNode(n)

	if a != nil {
		return HtmlNodeInfo{
			ParentTag:       a.Data,
			Content:         a.FirstChild.Data,
			QueryableString: CreateQueryableString(a),
		}, nil
	}

	return HtmlNodeInfo{}, errors.New("Could not find node")
}

func CheckPrice(c *colly.Collector, watcher entities.Watcher) (string, error) {
	price := ""

	tag, err := GetNodeQueryableHtmlWithParser(watcher.Node)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	c.OnHTML(tag.QueryableString, func(e *colly.HTMLElement) {
		if e.Text != "" {
			re := regexp.MustCompile("[0-9]+")
			numbers := (re.FindAllString(e.Text, -1))

			np := strings.Join(numbers, ".")

			f, err := strconv.ParseFloat(np, 64)

			if err == nil && f > 1 {
				price = np
			}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "\nError:", err)

		r.Request.Retry()
	})

	err = c.Visit(watcher.Url)

	if err != nil {
		return "", err
	}

	return price, nil
}
