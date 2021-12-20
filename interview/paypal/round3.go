package main

import (
	"fmt"
	"strings"
)



/*
input: stream char

if 80 char space, break,
if not break on space

input -> wordBreak -> output
*/


func wordBreak() func(rune) {
	count := 0
	return func(c rune) {
		count += 1
		if count >= 80 && c == ' ' {
			fmt.Printf("\n")
			count = 0
		}
	}
}

func main() {
	input := strings.Repeat("Hello ", 20)
	wBreak := wordBreak()

	for _, c := range input {
		wBreak(c)
		fmt.Printf("%c", c)
	}
}
