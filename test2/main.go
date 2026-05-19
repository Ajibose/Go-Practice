package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var colorMap = map[string]string{
	"black":  "\033[30m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"reset":  "\033[m",
}

func main() {
	args := os.Args
	// To ensure users provides just two arguments.
	if len(args) < 2 || len(args) > 4 {

		fmt.Println("Usage: go run . <text-to-draw>")
		return
	}

	color := flag.String("color", "", "Defines which color to use for the ascii")
	flag.Parse()
	args = flag.Args()

	var subString, input string
	input = args[0]
	if len(args) == 2 {
		subString = args[0]

		input = args[1]
	}

	if subString == "" {
		subString = input
	}

	// Making sure the program takes the second argument as input
	// input := args[0]

	// Empty string input must exit early, otherwise the banner drawing
	// logic would run on empty input and produce unexpected output.
	if input == "" {
		return
	}

	// This ensures that when \n is added to the
	// input, it transforms it to an actual new line
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Handles only an input of only a newline character
	if isOnlyNewLine(input) {
		for range input {
			fmt.Println()
		}
		return
	}

	// Split on newline because \n in input signals a new ASCII art row.
	words := strings.Split(input, "\n")

	contentLines := LoadBanner("standard")
	output := DrawAscii(words, contentLines, *color, subString)
	fmt.Print(output)

}

// isOnlyNewLine checks if the input contains only a newline or not.
func isOnlyNewLine(input string) bool {
	for _, ch := range input {
		if ch != '\n' {
			return false
		}
	}

	return true
}

// LoadBanner reads the banner file and returns a slice containing its lines
func LoadBanner(style string) []string {
	// The function accept only the style, extension should be added
	// to make it a complete filename
	fileName := fmt.Sprintf("%s.%s", style, "txt")

	content, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Printf("Error: Error reading file %v\n", err)
		os.Exit(1)
	}

	// To accomodate for the \r in window-based files
	modifiedContent := strings.ReplaceAll(string(content), "\r\n", "\n")

	contentLines := strings.Split(modifiedContent, "\n")

	return contentLines
}

// DrawAscii takes in two inputs and returns
// a string that has the contentlines to be drawn.
func DrawAscii(words []string, contentLines []string, color string, subString string) string {
	output := ""
	//substring := "som"
	//inside := false

	for _, word := range words {
		// A blank word results from splitting on \n, so emit a newline and move on.
		if word == "" {
			output += "\n"
			continue
		}

		for row := 0; row < 8; row++ {
			// for i, char := range word {
			for i := 0; i < len(word); i++ {
				// int(char)-32 gives the character offset from the first character(space)
				// 9 is multiplied because there are 8 lines of char + the newline
				// 1 is added because the first char(space) of the file doesn't start at line 1.
				start := 1 + (int(word[i])-32)*9
				if findNextSubString(i, word, subString) {
					k := 0
					output += colorMap[color]
					for k < len(subString) {
						start := 1 + (int(word[i])-32)*9
						output += contentLines[start+row]
						i++
						k++
					}
					output += colorMap["reset"]
					i--
					continue
				}

				output += contentLines[start+row]
			}
			output += "\n" // To ensure vertical alignment of lines
		}

	}
	return output
}

func findNextSubString(i int, s string, substring string) bool {

	if i+len(substring)-1 < len(s) && s[i:i+len(substring)] == substring {
		return true
	}

	return false
}
