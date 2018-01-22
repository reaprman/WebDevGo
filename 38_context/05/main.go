package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // make sure all paths cancel the context to avoid context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Second)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return // avoid leaking of this goroutine when ctx is done.
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}

/*
Unlike 04 this program exits and doesnt lose the goroutine so it exits.

Run Result:
0
1
2
3
4
5
*/
