package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

// this is important because if someone else tries to use the same key of the same type
// we may run into a colision
// so avoid we make our key of an unexported type
type key int

const requestIDKey key = 42

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("could not find ID in context")
		return
	}
	log.Printf("[%d] %v", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
