package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*Drill 4.5 — Non-capturing group
	GOAL: Match http://example.com or https://example.com, but do NOT capture the http/https part.
	INPUT: https://example.com
	Your FindStringSubmatch should return only one element (the whole match). If it returns two, you used the wrong group type. Look up (?:...). */
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`)

	inputs := []string{"2026-05-07T09:14:25", "2026-05-07", "09:14:25", "2026-05-07T9:14:25", "2026-05-07T09:14"}
	for _, in := range inputs {
		fmt.Printf("%-30q -> %v\n", in, re.MatchString(in))
	}
	/* input := "<b>hello</b>"
	greedy := regexp.MustCompile(`^<.*>$`)
	lazy := regexp.MustCompile(`^<.*?>$`)
	fmt.Printf("greedy: %q\n", greedy.FindString(input))
	fmt.Printf("lazy:   %q\n", lazy.FindString(input)) */
	/* re1 := regexp.MustCompile(`\d`)   // no anchors
	re2 := regexp.MustCompile(`^\d$`) // with anchors

	fmt.Println(re1.MatchString("abc5xyz")) // ? true
	fmt.Println(re1.MatchString("5"))       // ? true
	fmt.Println(re1.MatchString("abc"))		// ? false

	fmt.Println(re2.MatchString("abc5xyz")) // ? false
	fmt.Println(re2.MatchString("5"))      // ? true
	fmt.Println(re2.MatchString("abc"))     // ? false */

	/* chars := []string{"á", "ẹ", "ẹ́", "ẹ̀", "ọ́", "Olúwáṣẹ́gun"}
	for _, c := range chars {
		fmt.Printf("%-15q  bytes=%d  runes=%d\n",
			c, len(c), utf8.RuneCountInString(c))
	} */
}
