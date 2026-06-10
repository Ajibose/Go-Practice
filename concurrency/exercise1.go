package main

/*
Exercise 1 — Squares
Write a program that spawns 3 goroutines. Each goroutine receives a number as an argument, squares it, and sends the result back. Main prints all 3 results.
Numbers: 2, 3, 4
Expected output (any order):
4
9
16
*/

import "fmt"

func routine(n int, result chan<- int) {
	result <- n * n
}

func main() {
	result := make(chan int)

	go routine(1, result)
	go routine(2, result)
	go routine(3, result)

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
