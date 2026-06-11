package main

import "fmt"

type Bag map[string][]string

type Person struct {
	Name  string
	Stuff Bag
}

func main() {
	p := Person{
		Name: "Ibrahim",
		Stuff: Bag{
			"books": {"Go Programming"},
		},
	}

	fmt.Println(p.Stuff)
}
