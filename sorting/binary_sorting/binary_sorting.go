package main

import (
	"fmt"
)

// binarySearch finds the index where key should be inserted
func binarySearch(arr []int, item int, low int, high int) int {
	for low <= high {
		mid := low + (high-low)/2
		if item == arr[mid] {
			return mid + 1
		} else if item > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

// binaryInsertionSort sorts an array using binary search + insertion
func binaryInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		selected := arr[i]
		// Find location to insert using binary search
		loc := binarySearch(arr, selected, 0, i-1)

		// Shift elements to make room
		j := i - 1
		for j >= loc {
			arr[j+1] = arr[j]
			j--
		}
		// Insert at the correct location
		arr[loc] = selected
	}
}

func main() {
	arr := []int{4,9,0,2,8,7,1,4,4,7,3,2,1,5,6,5,-4}
	fmt.Println("Before sorting:", arr)

	binaryInsertionSort(arr)
	fmt.Println("After sorting: ", arr)
}