package main

import (
    "fmt"
    "reflect"
)

func one(xPtr *int) {
    *xPtr = 1
}

func main() {
    xPtr := new(int)
    one(xPtr)
    fmt.Println(reflect.TypeOf(xPtr), xPtr, *xPtr)
}
