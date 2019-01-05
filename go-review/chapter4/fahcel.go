package main

import "fmt"

func main () {

    var (
        fahdegree float64
        celcius float64
    )

    fmt.Printf("Enter temperature in F:")
    fmt.Scanf("%f", &fahdegree)

    celcius = ((fahdegree - 32.0) * 5/9)
    fmt.Println(celcius)

}
