package main

import "fmt"

func oddgen() func() int {
    i := -1
    return func() int {
        i += 2
        return i
    }
}

func main() {
    nextodd := oddgen()
    for i:= 0; i < 5; i++ {
        fmt.Println(nextodd())
    }
}
