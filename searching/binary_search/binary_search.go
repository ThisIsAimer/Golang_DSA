package main

import (
	"fmt"
	"math/rand"
)

func main() {
	OrginalArray := make([]int, 100)

	for i := range 100 {
		OrginalArray[i] = i + 1
	}
	myArray := OrginalArray

	// myValue := {any user given int}
	myValue := rand.Intn(99) + 1

	//to calculate the actual index
	indexOffset := 0

	operations := 1

	for {
		halfLen := len(myArray) / 2
		fmt.Println(halfLen)

		if myArray[halfLen] == myValue {
			fmt.Println("value found! no. of operations performed:", operations)
			fmt.Println("value needed is:", myValue)
			fmt.Println("value at index:", indexOffset+halfLen)
			fmt.Println("value at index is:", OrginalArray[indexOffset+halfLen])
			break
		} else if myValue > myArray[halfLen] {
			myArray = myArray[halfLen:]
			indexOffset += halfLen
		} else {
			myArray = myArray[:halfLen]
		}
		operations += 1
	}

}
