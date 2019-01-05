package main

import "fmt"

func main () {

    var x string
    x = "Hello World"
    fmt.Printf(x)
    fmt.Printf("\n")
    fmt.Printf(x + "!!!\n")


    y := "Infered as string"

    var z = "again infered as string with var statement"

    fmt.Println(y, z)


}
