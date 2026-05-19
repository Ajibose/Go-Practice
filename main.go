package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "hello World"
	res := ""

	newWord := true

	for _, ch := range str {
		if newWord {
			res += string(unicode.ToUpper(ch))
			newWord = false

		} else {
			newWord = true
			res += string(unicode.ToLower(ch))
		}

	}

	fmt.Println(res)
}
