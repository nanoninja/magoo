
package magoo

// Chain is a stack of Handlers.
type Chain struct {
	index    int8
	handlers []Handler
}

// NewChain returns a instance of Chain.
func NewChain() *Chain {
	c := &Chain{handlers: make([]Handler, 0)}
	c.Rewind()
	return c
}

// Count counts all handlers.
func (c *Chain) Count() int {
	return len(c.handlers)
}

// HasNext returns true if the iteration has more handler.
func (c *Chain) HasNext() bool {
	return c.index < int8(c.Count()-1)
}

// Index returns current slice index.
func (c *Chain) Index() int8 {
	return c.index
}

// Next moves to next handler.
func (c *Chain) Next() {
	c.index++
}

// Push appends a handler onto the handler stack.
func (c *Chain) Push(h Handler) {
	c.handlers = append(c.handlers, h)
}

// Rewind rewinds slice back to the start.
func (c *Chain) Rewind() {
	c.index = -1
}

// Value returns current slice handler.
func (c *Chain) Value() Handler {
	if c.Index() <= int8(-1) {
		return nil
	}
	return c.handlers[c.index]
}