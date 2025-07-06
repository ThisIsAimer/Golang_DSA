package main

import "fmt"

func main(){
	myArray := []int{4,9,0,2,8,7,1,4,4,7,3,2,1,5,6,5,-4}
	fmt.Println("default array:", myArray)
	noOfOperations := 0
	for{
		operations := 0
		for i := range myArray{
			if i == 0 {
				continue
			} else if myArray[i] < myArray[i-1]{
				tempVar := myArray[i]
				myArray[i] = myArray[i-1]
				myArray[i-1] = tempVar
				operations += 1
				noOfOperations += 1
			}
		}
		if operations == 0 {
			break
		}
	}



	fmt.Println("length of array:", len(myArray))
	fmt.Println("no. of operations performed:", noOfOperations)
	fmt.Println("sorted array is:", myArray)

}