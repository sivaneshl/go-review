package main

import "fmt"

func fibo() func() int {
    a,b := 1,0

    return func() int {
        a,b = b,a+b
        return a
    }

}

func main() {
    f := fibo()
    for i := 0; i < 8; i++ {
        fmt.Println(f())
    }
}
