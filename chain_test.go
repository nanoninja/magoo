package magoo

import "testing"

var getChainTests = []struct {
	handler Handler
}{
	{handler: HandlerFunc(func(c *Context) {})},
	{handler: HandlerFunc(func(c *Context) {})},
	{handler: HandlerFunc(func(c *Context) {})},
	{handler: HandlerFunc(func(c *Context) {})},
	{handler: HandlerFunc(func(c *Context) {})},
}

func TestNewChain(t *testing.T) {
	tests := getChainTests
	c := NewChain()

	for _, tt := range tests {
		c.Push(tt.handler)
	}
	if got, want := c.Count(), len(tests); got != want {
		t.Errorf("Count() got %d; want %d", got, want)
	}
	if got, want := c.Index(), int8(-1); got != want {
		t.Errorf("Index() got %d; want %d", got, want)
	}
	if got, want := c.HasNext(), true; got != want {
		t.Errorf("HasNext() got %v; want %v", got, want)
	}
	if got := c.Value(); got != nil {
		t.Errorf("Value() got %v; want nil", got)
	}
}

func TestChain(t *testing.T) {
	tests := getChainTests
	c := NewChain()

	for _, tt := range tests {
		c.Push(tt.handler)
	}

	c.Next()
	c.Next()

	if got, want := c.Index(), int8(1); got != want {
		t.Errorf("Index() got %d; want %d", got, want)
	}
	if got := c.Value(); got == nil {
		t.Errorf("Value() got %v; want Handler", got)
	}

	c.Rewind()

	if got, want := c.Index(), int8(-1); got != want {
		t.Errorf("Index() got %d; want %d", got, want)
	}
}