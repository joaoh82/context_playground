package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/joaoh82/context_playground/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	// this outputs value for foo is <nil>
	// And that is why you cannot use context to send values across servers
	// for that you need to use the request it self
	fmt.Printf("value for foo is %v\n", ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		// this is awesome, because if the request in canceled on the client side,
		// now we have way of knowing here in the server, and that meand that we could cancel any other operations
		// tighted to this request, so there is no waste of computer power
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
