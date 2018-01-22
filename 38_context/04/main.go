package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

// gen is a broken generator that will leak a goroutine
func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}

/*
Program never exits and keeps running due to leaked goroutine

Run Result:
0
1
2
3
4
5

*/
