package workers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// CrawlerWorker should run all functions about crawler
type CrawlerWorker struct {
	Worker
}

// Notify should receive the messages
func (c *CrawlerWorker) Notify(message string) {
	if c == nil {
		return
	}
	if message == "startup" {
		c.parseHTML("https://google.com")
	}
	fmt.Println("CrawlerWorker get message:", message)
}

// Send should send the messages
func (c *CrawlerWorker) Send(message string) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}

// NewCrawlerWorker returns a instance from worker
func NewCrawlerWorker(mediator IMediator) *CrawlerWorker {
	return &CrawlerWorker{Worker{mediator}}
}

func (c *CrawlerWorker) mapElements(s *goquery.Selection) {
	className, exists := s.Attr("class")
	if exists {
		c.Send(className)
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
