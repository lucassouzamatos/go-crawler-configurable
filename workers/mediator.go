package workers

// IMediator is a interface for god
type IMediator interface {
	Send(WrapperMessage, IWorker)
}

// IWorker is a interface for workers
type IWorker interface {
	Send(WrapperMessage)
	Notify(WrapperMessage)
}

// Mediator is a constructor
type Mediator struct {
}

// Worker is a constructor
type Worker struct {
	mediator IMediator
}
