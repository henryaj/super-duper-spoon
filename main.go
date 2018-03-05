package main

import (
	"io"
	"log"
	"net/http"
)

//Cache is a generic cache interface.
type Cache interface {
	setHandler(http.ResponseWriter, *http.Request)
	getHandler(http.ResponseWriter, *http.Request)
}

//InMemoryCache is am in-memory implementation of Cache.
type InMemoryCache struct {
	keys map[string]string
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		keys: make(map[string]string),
	}
}

func (c *InMemoryCache) setHandler(writer http.ResponseWriter, req *http.Request) {
	log.Println("setting...")

	params := req.URL.Query()
	for k, v := range params {
		c.keys[k] = v[len(v)-1] // save only the last param for a given key
	}

	log.Println("Cache now contains:", c.keys)
}

func (c *InMemoryCache) getHandler(writer http.ResponseWriter, req *http.Request) {
	log.Println("getting...")

	param := req.URL.Query().Get("key")

	v, ok := c.keys[param]
	if ok {
		io.WriteString(writer, v)
	}
}

func main() {
	c := NewInMemoryCache()

	http.HandleFunc("/set", c.setHandler)
	http.HandleFunc("/get", c.getHandler)

	log.Println("Started")

	log.Fatal(http.ListenAndServe(":4000", nil))
}
