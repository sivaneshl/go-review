package main

import "fmt"

func fibo(n int) []int {

    a, b := 0, 1
    c := a + b
    ser := make([]int, n)

    for i := 0; i < n; i++{
        if i == 0 || i == 1 {
            ser[i] = i
        } else {
            ser[i] = c
            a = b
            b = c
            c = a + b
        }
    }

    return (ser)

}

func main() {
    fmt.Println(fibo(7))
}
