package magoo

// A Param represents the key-value pairs in an Magoo context.
type Param map[string]string

// Add adds the key, value pair to the header.
// It appends to any existing values associated with key.
func (p Param) Add(key, value string) {
	if _, ok := p[key]; ok {
		return
	}
	p[key] = value
}

// Del deletes the values associated with key.
func (p Param) Del(key string) {
	delete(p, key)
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns "".
func (p Param) Get(key string) string {
	if v, ok := p[key]; ok {
		return v
	}
	return ""
}

// Set sets the header entries associated with key to
// the single element value. It replaces any existing
// values associated with key.
func (p Param) Set(key, value string) {
	p[key] = value
}
