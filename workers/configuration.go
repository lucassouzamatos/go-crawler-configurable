package workers

// HandlerConfiguration should store the configuration about specific website
type HandlerConfiguration struct {
	URL string
}

// Configuration should store the configuration about all crawler
type Configuration struct {
	Handlers []HandlerConfiguration
}
