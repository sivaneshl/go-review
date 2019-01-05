package main

import "fmt"

func main () {

    const oneft = 0.3048
    var (
        foot float64
        meter float64
    )

    fmt.Printf("Enter distance in Ft: ")
    fmt.Scanf("%f", &foot)

    meter = oneft * foot
    fmt.Println(foot, "Ft =", meter,"m")

}
