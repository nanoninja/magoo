package magoo

import "net/http"

type Handler interface {
	ServeHTTP(c *Context)
}

type HandlerFunc func(c *Context)

func (f HandlerFunc) ServeHTTP(c *Context) {
	f(c)
}

func adapt(h http.Handler) Handler {
	return HandlerFunc(func(c *Context) {
		h.ServeHTTP(c.ResponseWriter, c.Request)
		c.Next()
	})
}