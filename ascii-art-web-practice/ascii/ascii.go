package ascii

import (
	"strings"
)

func GenerateAsciiArt(style string, text string) string {

	var output strings.Builder

	ReplacedContent := strings.ReplaceAll(style, "\r\n", "\n")
	contentLines := strings.Split(ReplacedContent, "\n")

	for row := 0; row < 8; row++ {
		for _, ch := range text {
			start := 1 + ((int(ch) - 32) * 9)
			output.WriteString(contentLines[start+row])
		}
		output.WriteString("\n")
	}

	return output.String()
}
