# koala [![Build Status](https://travis-ci.org/haochi/koala.svg?branch=master)](https://travis-ci.org/haochi/koala)

lightweight multiplexer for Go

## install

```bash
$ go get "github.com/haochi/koala"
```

```go
package main

import (
	"github.com/haochi/koala"
	"net/http"
	"fmt"
)

func main() {
	app := koala.New()

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	app.Get("/~{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user %s reporting for duty!", koala.Param(r, "id"))
	})

	panic(http.ListenAndServe(":8080", app))
}

```

## documentation

[it is on godoc.org](https://godoc.org/github.com/haochi/koala)
