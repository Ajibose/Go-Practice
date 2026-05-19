package main

import "fmt"

func buildGreeters(names []string) []func() string {
	greeters := []func() string{}
	for _, name := range names {
		greeter := func() string {
			return "Hello, " + name + "!"
		}
		greeters = append(greeters, greeter)
	}
	return greeters
}

func main() {
	names := []string{"Ada", "Bob", "Cho"}
	greeters := buildGreeters(names)
	for _, g := range greeters {
		fmt.Println(g())
	}
}