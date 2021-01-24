package workers

import (
	"fmt"
)

// CloudVisionWorker should run all functions about cloud vision
type CloudVisionWorker struct {
	Worker
}

// Notify should receive the messages
func (c *CloudVisionWorker) Notify(message WrapperMessage) {
	if c == nil {
		return
	}
	if message.text == "mapped-element" {
		c.Extract(message)
	}
	fmt.Println("CloudVisionWorker get message:", message.text)
}

// Send should send the messages
func (c *CloudVisionWorker) Send(message WrapperMessage) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}

// NewCloudVisionWorker returns a instance from worker
func NewCloudVisionWorker(mediator IMediator) *CloudVisionWorker {
	return &CloudVisionWorker{Worker{mediator}}
}

func (c *CloudVisionWorker) Extract(message WrapperMessage) {
	println("extracting message")
}
