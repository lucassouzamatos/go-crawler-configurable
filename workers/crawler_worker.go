package workers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/net/html"
)

// CrawlerWorker should run all functions about crawler
type CrawlerWorker struct {
	Worker
}

// Notify should receive the messages
func (c *CrawlerWorker) Notify(message WrapperMessage) {
	if c == nil {
		return
	}
	if message.text == "startup" && message.configuration != nil {
		fmt.Println("CrawlerWorker get configuration url:", message.configuration.URL)
		c.parseHTML(message.configuration.URL)
	}

	fmt.Println("CrawlerWorker get message:", message.text)
}

// Send should send the messages
func (c *CrawlerWorker) Send(message WrapperMessage) {
	if c == nil {
		return
	}
	// 	c.parseHTML(url)
	c.mediator.Send(message, c)
}

// NewCrawlerWorker returns a instance from worker
func NewCrawlerWorker(mediator IMediator) *CrawlerWorker {
	return &CrawlerWorker{Worker{mediator}}
}

func (c *CrawlerWorker) mapElements(s *goquery.Selection) {
	_, exists := s.Attr("class")
	if exists {
		var u, err = uuid.NewV4()
		if err == nil {
			ws := WrapperSelection{selection: s, identifier: u}
			e := WrapperMessage{text: "mapped-element", querySelection: &ws}

			c.Send(e)
		}
	}

	if s.Children() != nil {
		s.Children().Each(func(_ int, ch *goquery.Selection) {
			c.mapElements(ch)
		})
	}
}

func (c *CrawlerWorker) parseHTML(url string) {
	resp, reqErr := http.Get(url)
	if reqErr != nil {
		fmt.Println(reqErr)
	}
	defer resp.Body.Close()

	root, _ := html.Parse(resp.Body)

	doc := goquery.NewDocumentFromNode(root)
	doc.Each(func(_ int, s *goquery.Selection) {
		c.mapElements(s)
	})
}
