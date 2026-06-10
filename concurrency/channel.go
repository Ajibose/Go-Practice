package main

import "fmt"

func main() {
	ch := make(chan int)

	func() {
		ch <- 42
	}()

	value := <-ch
	fmt.Println(value)
}
