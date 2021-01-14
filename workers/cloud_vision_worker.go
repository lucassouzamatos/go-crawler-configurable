package workers

import "fmt"

// CloudVisionWorker should run all functions about cloud vision
type CloudVisionWorker struct {
	Worker
}

// Notify should receive the messages
func (c *CloudVisionWorker) Notify(message string) {
	if c == nil {
		return
	}
	fmt.Println("CloudVisionWorker get message:", message)
}

// Send should send the messages
func (c *CloudVisionWorker) Send(message string) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}
