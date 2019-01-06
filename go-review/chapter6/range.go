package main

import "fmt"

func main () {

    nums := []int{10,20,30,40,50}

    for i, value := range(nums){
        fmt.Println(i, value)
    }
}
