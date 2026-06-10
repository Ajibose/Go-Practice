package main

import (
	"fmt"
	"time"
)

/*
Exercise 3 — Select with timeout
Write a program where a goroutine sends a value after 300ms. Main should wait for it but give up after 200ms and print "timed out" instead.
Hint: time.After returns a channel.
*/

func main() {
	ch := make(chan string)
	t := time.After(200 * time.Millisecond)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- "Waiting for you"
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
			return
		case <-t:
			fmt.Println("timed out")
			return
		}
	}
}
