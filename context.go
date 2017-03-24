package magoo

import (
	"net"
	"net/http"
	"strings"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Param          Param
	chain          *Chain
}

func (c *Context) ClientIP() string {
	for _, addr := range []string{"X-Forwarded-For", "X-Real-IP"} {
		addr = c.Request.Header.Get(addr)
		if ip := strings.TrimSpace(addr); len(ip) > 0 {
			return ip
		}
	}
	if ip := strings.TrimSpace(c.Request.RemoteAddr); len(ip) > 0 {
		ip, _, _ = net.SplitHostPort(ip)
		return ip
	}
	return ""
}

func (c *Context) IsMethod(method string) bool {
	return c.Request.Method == method
}

// IsSecure returns true if the request scheme is https.
func (c *Context) IsSecure() bool {
	return c.Request.TLS != nil
}

func (c *Context) IsXMLHTTPRequest() bool {
	return c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

func (c *Context) Next() {
	if c.chain.HasNext() {
		c.chain.Next()
		c.chain.Value().ServeHTTP(c)
	}
}