package main

import "fmt"

func main () {

    var sli1 []float64
    fmt.Println(sli1)

    sli2 := make([]float64, 5)
    fmt.Println(sli2)

    sli3 := make([]float64, 5, 10)
    fmt.Println(sli3)

    arr := [5]float64{2,4,6,8,10}
    sli4 := arr[1:4]
    fmt.Println(sli4)

    x := [6]string{"a","b","c","d","e","f"}
    fmt.Println(x[2:5])

}
