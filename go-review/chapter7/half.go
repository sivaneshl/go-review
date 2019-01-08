package main

import "fmt"

func half(x int) (halfval int, iseven bool) {
    halfval = x / 2
    if x % 2 == 0 {
        iseven = true
    } else {
        iseven = false
    }
    return
}

func main() {
    fmt.Println(half(5))
    fmt.Println(half(8))
}
