package magoo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdapte(t *testing.T) {
	handler := adapt(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(415)
	}))
	w := httptest.NewRecorder()
	c := &Context{
		ResponseWriter: w,
		Request:        (*http.Request)(nil),
		Param:          Param{},
		chain:          NewChain(),
	}
	handler.ServeHTTP(c)

	if got, want := w.Code, 415; got != want {
		t.Errorf("Code got %d; want %d", got, want)
	}
}