package main

import (
	"fmt"
)

func main() {
	strs := []string{"cat", "elephant", "dog", "lion", "tiger"}

	n := len(strs)

	// Bubble sort by length
	for i := 0; i < n-1; i++ {
		operations := 0
		for j := 0; j < n-i-1; j++ {
			if len(strs[j]) > len(strs[j+1]) {
				strs[j], strs[j + 1] = strs[j + 1], strs[j]
				operations++
			}
		}
		if operations == 0 {
			break
		}
	}

	fmt.Println("Sorted by length (ascending):", strs)
}
