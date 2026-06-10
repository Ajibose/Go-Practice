package main

import "fmt"

/*
Exercise 2 — Range and close
Write a function generate that takes a slice of integers, sends each one into a channel, closes it when done, and returns the channel:
gofunc generate(nums []int) <-chan int
In main, use a range loop to print every value that comes out.
*/

func generate(nums []int) <-chan int {
	ch := make(chan int)

	go func() {
		for _, num := range nums {
			ch <- num
		}

		close(ch)
	}()

	return ch
}

func main() {
	c := generate([]int{1, 2, 3, 4, 5})
	for v := range c {
		fmt.Println(v)
	}

}
