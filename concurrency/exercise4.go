package main

import "fmt"

/*
Exercise 4 — Pipeline
Write a function square that takes a receive-only channel, squares each value it receives, and returns a new channel with the results:
gofunc square(in <-chan int) <-chan int
In main, create a channel, send 3 into it, then chain 3 calls to square:
go// expected: 6561
fmt.Println(<-square(square(square(c))))
Each square call should spawn its own goroutine internally — same pattern as generate from Exercise 2.
*/

func square(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for v := range in {
			ch <- v * v
		}
	}()

	return ch
}

func main() {
	c := make(chan int)

	out := square(square(square(c)))
	c <- 3

	fmt.Println(<-out)
}
