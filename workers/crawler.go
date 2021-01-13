package workers

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type CrawlerWorker struct {
}

// NewCrawlerWorker should creates a new worker
func NewCrawlerWorker() *CrawlerWorker {
	return &CrawlerWorker{}
}

// MapElements should walks all children elements recursively
func (crawlerWorker *CrawlerWorker) MapElements(s *goquery.Selection) {
	// class, exists := s.Attr("class")

	if s.Children() != nil {
		s.Children().Each(func(_ int, c *goquery.Selection) {
			crawlerWorker.MapElements(c)
		})
	}
}

// ParseHTML should returns a HTML node from uri
func (crawlerWorker *CrawlerWorker) ParseHTML(url string) {
	resp, reqErr := http.Get(url)
	if reqErr != nil {
		fmt.Println(reqErr)
	}
	defer resp.Body.Close()

	root, _ := html.Parse(resp.Body)

	doc := goquery.NewDocumentFromNode(root)
	doc.Each(func(_ int, s *goquery.Selection) {
		crawlerWorker.MapElements(s)
	})
}
