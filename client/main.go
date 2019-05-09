package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, time.Second)
	// defer cancel()

	// this outputs value for foo is <nil>
	// And that is why you cannot use context to send values across servers
	// for that you need to use the request it self
	ctx = context.WithValue(ctx, "foo", "bar")

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Why do we need to do this? Why isn't there a context field inside the request struct it self.
	// Mainly because of backwards compatibility reasons with Go 1.0, that did have the Context package yet.
	req = req.WithContext(ctx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}
	io.Copy(os.Stdout, res.Body)
}
