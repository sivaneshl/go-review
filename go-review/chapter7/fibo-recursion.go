package main

import "fmt"

func fibo(x int) int {
    if x == 0 {
        return 0
    }
    if x == 1 {
        return 1
    }
    return fibo(x-1) + fibo(x-2)
}

func main() {
    sum := 0
    for i:=1; i<=7; i++ {
        sum = sum + fibo(i)
        // fmt.Println(i,sum)
    }
    fmt.Println(sum)
}
