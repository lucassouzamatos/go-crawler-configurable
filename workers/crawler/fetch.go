package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// MapElements should walks all children elements recursively
func MapElements(s *goquery.Selection) {
	// class, exists := s.Attr("class")

	if s.Children() != nil {
		s.Children().Each(func(_ int, c *goquery.Selection) {
			MapElements(c)
		})
	}
}

// ParseHTML should returns a HTML node from uri
func ParseHTML(url string) {
	resp, reqErr := http.Get(url)
	if reqErr != nil {
		fmt.Println(reqErr)
	}
	defer resp.Body.Close()

	root, _ := html.Parse(resp.Body)

	doc := goquery.NewDocumentFromNode(root)
	doc.Each(func(_ int, s *goquery.Selection) {
		MapElements(s)
	})
}

func main() {
	ParseHTML("https://pt.wikipedia.org/wiki/Wikip%C3%A9dia:P%C3%A1gina_principal")
}
