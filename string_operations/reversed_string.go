package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Convert string to a slice of runes to handle Unicode characters
	runes := []rune(s)
	n := len(runes)

	// Swap characters from both ends
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	str := "Hello, my man"
	reversed := reverseString(str)
	fmt.Println("Original:", str)
	fmt.Println("Reversed:", reversed)
}