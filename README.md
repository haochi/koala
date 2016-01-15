# koala

lightweight multiplexer for Go

## install

```bash
$ go get "github.com/haochi/koala"
```

```golang
import "github.com/haochi/koala"

app := koala.New()

app.Get("/~{id}", func (writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "user %s reporting for duty!", koala.GetPathParam(request, "id"))
})

app.ListenAndServe(":8080")
```
