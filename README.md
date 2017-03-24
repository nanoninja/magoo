# Magoo

Magoo is experimental HTTP Middleware handler for Go..

# Basic example

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/nanoninja/magoo"
)

func main() {
	m := magoo.New()
	m.Use(Log)
	m.Use(Auth)
	m.Use(Home)

	fmt.Println("Server is running...")
	http.ListenAndServe(":3000", m)
}

func Auth(c *magoo.Context) {
	fmt.Println("Auth Before")
	c.Next()
	fmt.Println("Auth After")
}

func Log(c *magoo.Context) {
	fmt.Println("Log Before")
	c.Next()
	fmt.Println("Log After")
}

func Home(c *magoo.Context) {
	fmt.Println("Homepage")
	c.ResponseWriter.Write([]byte("Homepage"))
}
```