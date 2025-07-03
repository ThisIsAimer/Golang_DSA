package main

import "fmt"

func main(){
    myArray := []int{4,9,0,2,8,7,1}

    a := 0

    for _ , v := range myArray{
        if a <= v{
            a = v
        }
    }

    fmt.Println("largest value:", a)

    b := 0

    for _ , v := range myArray{
        if v != a && b <= v {
            b = v
        }
    }
	
    fmt.Println("second largest value:", b)

}