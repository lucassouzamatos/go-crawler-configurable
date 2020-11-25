package handlers

import (
	"net/http"
)

// Configuration is the structure for the process execution
type Configuration struct {}

// DefineConfiguration is a function API
func DefineConfiguration(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from go worker"))
}