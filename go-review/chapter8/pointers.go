package main

import (
    "fmt"
    "reflect"
)

func zero(xPtr *int) {
    *xPtr = 0
    // fmt.Println(reflect.TypeOf(xPtr))

}

func main() {
    x := 5
    zero(&x)
    fmt.Println(x)
}
