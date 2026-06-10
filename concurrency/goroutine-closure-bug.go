package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
		go func() {
			fmt.Println(i) // which i does this capture?
		}()
	}

	time.Sleep(5 * time.Second)
}

// My answer - Wrong
/*
4, because by the time the goroutines get popped from the run queue,
the loop would have already finished, so i will be 4.
The main goroutine would have also exited so no goroutine gets executed but they capture i as 4(without Sleep)
*/

//Correction
/*
Go loop variable capture behavior changed in 1.22. Code that demonstrates the closure bug on older versions may behave correctly on 1.22+.

In version 1.22+, each goroutine create their own copy of the variable.
*/