package main

import (
	"fmt"
	"strings"
)

func longestPalindromicSubstring(s string) string {
	if len(s) <= 1 {
		return s
	}

	// Preprocess the input string with special characters
	preprocessed := preprocessString(s)
	fmt.Println(preprocessed)

	// Initialize an array to store the lengths of palindromes centered at each position
	p := make([]int, len(preprocessed))

	// Initialize variables to keep track of the center of the longest palindrome found so far
	center, maxRight := 0, 0

	// Iterate through the preprocessed string to compute palindrome lengths
	for i := 1; i < len(preprocessed)-1; i++ {
		// Calculate the mirror index
		mirror := 2*center - i

		// Update the length of the palindrome centered at position i
		if i < maxRight {
			p[i] = min(maxRight-i, p[mirror])
		}

		// Attempt to expand the palindrome centered at position i
		for preprocessed[i+p[i]+1] == preprocessed[i-p[i]-1] {
			p[i]++
		}

		// Update the center and maxRight if necessary
		if i+p[i] > maxRight {
			center = i
			maxRight = i + p[i]
		}
	}

	// Find the maximum palindrome length and its center
	maxLength, maxCenter := 0, 0
	for i := 1; i < len(preprocessed)-1; i++ {
		if p[i] > maxLength {
			maxLength = p[i]
			maxCenter = i
		}
	}

	// Extract the longest palindromic substring from the preprocessed string
	start := (maxCenter - maxLength) / 2
	end := start + maxLength

	return s[start:end]
}

func preprocessString(s string) string {
	preprocessed := make([]byte, 0, 2*len(s)+3)
	var sb strings.Builder
	sb.WriteString("^")
	for _, char := range s {
		sb.WriteString(fmt.Sprintf("#%c", char))
	}
	preprocessed = append(preprocessed, '#', '$')
	sb.WriteString("#$")
	return sb.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindromicSubstring("babad"))  // Output: "bab" or "aba"
	fmt.Println(longestPalindromicSubstring("cbbd"))   // Output: "bb"
	fmt.Println(longestPalindromicSubstring("aaamma")) // Output: "bb"
}
