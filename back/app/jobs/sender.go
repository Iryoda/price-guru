package jobs

import (
	"context"
	"time"

	"github.com/iryoda/price-guru/app/entities"
	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
)

type Sender struct{}

func (s Sender) Send() {
	InitDotEnv()

	c := InitMongo()
	ch := InitAmqp()

	d := c.Database("price-guru").Collection("watchers")

	now := time.Now()
	var watchers []entities.Watcher

	cursor, err := d.Find(context.Background(), bson.M{})

	if err != nil {
		panic(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var watcher entities.Watcher
		err := cursor.Decode(&watcher)

		if err != nil {
			panic(err)
		}

		if watcher.Status == entities.STARTED || now.Sub(watcher.LastRun).Hours() > 24 {
			watchers = append(watchers, watcher)
		}
	}

	d.UpdateMany(context.Background(), bson.M{"_id": bson.M{"$in": watchers}}, bson.M{"$set": bson.M{"status": entities.SCHEDULED}})

	for _, watcher := range watchers {
		ch.PublishWithContext(context.TODO(),
			"watchers",
			"",
			false,
			false,
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(watcher.Id),
			},
		)
	}
}
