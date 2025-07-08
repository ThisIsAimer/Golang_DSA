package main

import "fmt"

func main() {

	star := 1

    for i := 5; i >= 0; i-- {

        for j := 0 ; j <= i ; j++ {
            fmt.Print(" ")
        }

		for range star{
			fmt.Print("*")
		}
		
        fmt.Print("\n")
		star += 2
    }

}