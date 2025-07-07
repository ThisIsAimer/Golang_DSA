package main

import "fmt"

func main(){

	myString := ""

	for i := 0; i <= 5 ; i++{
		for j := 0; j <= i ; j++{
			myString += "*"
		}
		myString += "\n"
	}
	fmt.Println("star pattern:")
	fmt.Print(myString)

}