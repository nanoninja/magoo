package magoo

import "testing"

func TestParam(t *testing.T) {
	p := &Param{"foo": "bar"}

	if got, want := p.Get("test"), ""; got != want {
		t.Errorf("Get() got %q; want %q", got, want)
	}
	if got, want := p.Get("foo"), "bar"; got != want {
		t.Errorf("Get() got %q; want %q", got, want)
	}

	p.Set("name", "gopher")
	if got, want := p.Get("name"), "gopher"; got != want {
		t.Errorf("Get() got %q; want %q", got, want)
	}

	p.Add("name", "golang")
	if got, want := p.Get("name"), "gopher"; got != want {
		t.Errorf("Get() got %q; want %q", got, want)
	}

	p.Add("john", "doe")
	if got, want := p.Get("john"), "doe"; got != want {
		t.Errorf("Get() got %q; want %q", got, want)
	}

	p.Set("test", "go")
	p.Del("test")
	if got, want := p.Get("test"), ""; got != want {
		t.Errorf("Get() got %q; want %q", got, want)
	}
}