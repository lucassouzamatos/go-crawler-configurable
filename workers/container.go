package workers

import (
	"sync"
)

type Container struct {
	mux      sync.RWMutex
	services map[string]interface{}
}

// Add service
func (c *Container) Add(name string, object interface{}) {
	if c.services == nil {
		c.services = make(map[string]interface{})
	}
	c.mux.Lock()
	c.services[name] = object
	c.mux.Unlock()
}

// Remove service
func (c *Container) Remove(name string) *Container {
	c.mux.Lock()
	delete(c.services, name)
	c.mux.Unlock()

	return c
}

// Get a service
func (c *Container) Get(name string) (object interface{}, ok bool) {
	c.mux.RLock()
	object, ok = c.services[name]
	c.mux.RUnlock()
	return object, ok
}
