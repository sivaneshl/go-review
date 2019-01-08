package main

import "fmt"

func swap(x,y *int) {
    z := 0
    z = *x
    *x = *y
    *y = z
}

func main() {
    x,y := 5,10

    swap(&x,&y)
    fmt.Println(x,y)
}
