package magoo

import "net/http"

// Magoo is a stack of Handlers that can be invoked as an http.Handler.
type Magoo struct {
	chain *Chain
}

// New returns a instance of Magoo.
func New(handlers ...Handler) *Magoo {
	m := &Magoo{chain: NewChain()}
	for _, h := range handlers {
		m.Use(h)
	}
	return m
}

// Use Uses the specified middleware handler.
func (m *Magoo) Use(i interface{}) {
	var h Handler
	switch v := i.(type) {
	case Handler, HandlerFunc:
		h = v.(Handler)
	case func(c *Context):
		h = HandlerFunc(v)
	case http.Handler, http.HandlerFunc:
		h = adapt(v.(http.Handler))
	case func(http.ResponseWriter, *http.Request):
		h = adapt(http.HandlerFunc(v))
	default:
		h = nil
	}
	if h != nil {
		m.chain.Push(h)
	}
}

func (m *Magoo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		ResponseWriter: w,
		Request:        r,
		Param:          Param{},
		chain:          m.chain,
	}
	ctx.Next()
	ctx.chain.Rewind()
}