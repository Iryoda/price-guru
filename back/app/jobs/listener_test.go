package jobs

import (
	"testing"

	"github.com/gocolly/colly"
	"github.com/imroc/req/v3"
	"github.com/iryoda/price-guru/app/entities"
)

func TestHtmlParser(t *testing.T) {
	tag := `<p id="valVista" class="val-prod valVista">R$ 2.139,90</p>`

	expected := HtmlNodeInfo{
		ParentTag:       "p",
		Content:         "R$ 2.139,90",
		QueryableString: `p[id='valVista'][class='val-prod valVista']`,
	}

	r, err := GetNodeQueryableHtmlWithParser(tag)

	if err != nil {
		t.Errorf(err.Error())
	}

	if r != expected {
		t.Errorf("Expected %s, got %s", expected, r)
	}

}

func TestHtmlParserDiv(t *testing.T) {
	tag := `<div class="jss290">R$ 1.949,99</div>`
	expected := HtmlNodeInfo{
		ParentTag:       "div",
		Content:         "R$ 1.949,99",
		QueryableString: `div[class='jss290']`,
	}

	r, err := GetNodeQueryableHtmlWithParser(tag)

	if err != nil {
		t.Errorf(err.Error())
	}

	if r != expected {
		t.Errorf("Expected %s, got %s", expected, r)
	}
}

func TestCrawler(t *testing.T) {
	w := entities.Watcher{
		Url:       "https://www.terabyteshop.com.br/produto/26672/kit-fan-lian-li-com-3-unidades-uni-fan-al120-v2-120mm-argb-black-uf-al120v2-3b",
		Node:      `<p id="valVista" class="val-prod valVista">R$ 587,90</p>`,
		LastValue: 587.90,
	}

	expected := "587.90"

	fakeChrome := req.DefaultClient().ImpersonateChrome()
	cc := colly.NewCollector(
		colly.UserAgent(fakeChrome.Headers.Get("User-Agent")),
	)
	cc.WithTransport(fakeChrome.Transport)

	res, err := CheckPrice(cc, w)

	if err != nil {
		t.Errorf(err.Error())
	}

	if expected != res {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}
