package workers

import (
	"encoding/json"
	"fmt"
	"os"
)

// CrawlerMediator is observable struct
type CrawlerMediator struct {
	Mediator
	workers []IWorker
}

// AddWorker append the worker in mediator list workers
func (m *CrawlerMediator) AddWorker(w IWorker) {
	if m == nil {
		return
	}
	m.workers = append(m.workers, w)
}

// Send should notify all the workers
func (m *CrawlerMediator) Send(message string, c IWorker) {
	if m == nil {
		return
	}
	for _, val := range m.workers {
		if c == val {
			continue
		}
		val.Notify(message)
	}
}

// SendAll should notify all the workers from the mediator
func (m *CrawlerMediator) SendAll(message string) {
	if m == nil {
		return
	}
	for _, val := range m.workers {
		val.Notify(message)
	}
}

// NewCrawler should creates a new worker
func NewCrawler() *CrawlerMediator {
	c := &CrawlerMediator{}
	c.AddWorker(NewCrawlerWorker(c))
	return c
}

// Start is the function that start execution
func (m *CrawlerMediator) Start() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
	}

	for _, handle := range configuration.Handlers {
		println(handle.URL)
	}
}
