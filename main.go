package main

import "./workers"

func main() {
	crawler := workers.NewCrawler()
	crawler.Start()
}
