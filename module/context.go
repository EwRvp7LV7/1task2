package module

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(400 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func operation3(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation3")
	}
}

func Context1() {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context
	// ctx, cancel := context.WithCancel(ctx)

	// The context will be cancelled after 3 seconds
	// If it needs to be cancelled earlier, the `cancel` function can
	// be used, like before
	//если tms больше 400 из operation1, то  operation1 отменит cancel()-ом
	//иначе по таймауту
	tms := 500
	ctx, cancel := context.WithTimeout(ctx, time.Duration(tms)*time.Millisecond)

	// // Setting a context deadline is similar to setting a timeout, except
	// // you specify a time when you want the context to cancel, rather than a duration.
	// // Here, the context will be cancelled on 2009-11-10 23:00:00
	// ctx, cancel := context.WithDeadline(ctx, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))

	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			fmt.Println("cancel context")
			cancel() //`cancel` function!
		}
	}()

	// Run operation2 with the same context we use for operation1
	go operation2(ctx)
	go operation3(ctx)

	<-time.After(500 * time.Millisecond)
}

//=================================================================
// we need to set a key that tells us where the data is stored
const keyID = "id"

func operation11(ctx context.Context) {
	// do some work

	// we can get the value from the context by passing in the key
	log.Println("operation1 for id:", ctx.Value(keyID), " completed")
	operation12(ctx)
}

func operation12(ctx context.Context) {
	// do some work

	// this way, the same ID is passed from one function call to the next
	log.Println("operation2 for id:", ctx.Value(keyID), " completed")
}

func ContextPassValue() {
	rand.Seed(time.Now().Unix())
	ctx := context.WithValue(context.Background(), keyID, rand.Int())
	operation11(ctx)
}

//=================================================================
