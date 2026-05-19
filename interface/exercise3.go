package inter

import "strconv"

func boolToString(c bool) string {
	if c {
		return "True"
	}
	return "False"
}

func FormatData(args ...interface{}) string {
	resultStr := ""
	for _, i := range args {
		switch t := i.(type) {
		case int:
			resultStr += "Number: "
			resultStr += strconv.Itoa(t)
			resultStr += ", "
		case string:
			resultStr += "Text: "
			resultStr += t
			resultStr += ", "
		case bool:
			resultStr += "Condition: "
			resultStr += boolToString(t)
			resultStr += ", "
		default:
			resultStr += "Unknown Type, "
		}
	}

	return resultStr[:len(resultStr)-2]
}


package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := os.ReadFile("standard.txt")
	check(err)

	str := "hello world"

	lines := strings.Split(string(data), "\n")

	for _, ch := range str {
		code := int(ch - 32)
		lineNumber := (code - 32) * 9
		character := lines[lineNumber : lineNumber+8]

		for _, line := range character {

			fmt.Println(line)
		}

	}
}
