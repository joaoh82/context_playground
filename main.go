package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	// creates the parent context with no values or cancelations
	ctx := context.Background()

	// WithCancel returns a copy of parent with a new Done channel.
	// The returned context's Done channel is closed when the returned cancel function is called or
	// when the parent context's Done channel is closed, whichever happens first.
	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)
	// the code below does basically the same as the line above. It's just a bit more explicit
	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
	// Canceling this context releases resources associated with it,
	// so code should call cancel as soon as the operations running in this
	// That is why we have the defer cancel() at the end, because even with the timeout set, we need to call cancel()
	// Also this code does basically the same the snippets above
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// sleepAndTalk(ctx, 5*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, wait time.Duration, message string) {
	select {
	case <-time.After(wait):
		fmt.Println(message)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
