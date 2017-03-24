package magoo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMagoo(t *testing.T) {
	m := New(nil)
	if got, want := m.chain.Count(), 0; got != want {
		t.Errorf("chain.Count() got %d; want %d", got, want)
	}
	m = New(
		HandlerFunc(func(c *Context) {}),
		HandlerFunc(func(c *Context) {}),
	)
	if got, want := m.chain.Count(), 2; got != want {
		t.Errorf("chain.Count() got %d; want %d", got, want)
	}
}

type testHandler struct{}

func (h *testHandler) ServerHTTP(c *Context) {}

func TestMagoo(t *testing.T) {
	test := []struct {
		handler interface{}
	}{
		{&testHandler{}},
		{HandlerFunc(func(c *Context) {})},
		{func(c *Context) {}},
		{http.NewServeMux()},
		{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{func(w http.ResponseWriter, r *http.Request) {}},
	}
	m := New()
	for _, tt := range test {
		m.Use(tt.handler)

		w := httptest.NewRecorder()
		m.ServeHTTP(w, (*http.Request)(nil))

		if got, want := w.Code, 200; got != want {
			t.Errorf("Code got %d; want %d", got, want)
		}
	}
}