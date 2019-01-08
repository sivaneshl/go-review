package main

import ( "fmt"; "math" )

type Circle struct {
    x, y, r float64
}

func circleArea(c Circle) float64 {
    return math.Pi * c.r * c.r
}

func main() {
    c1 := new(Circle)
    c1.x, c1.y, c1.r = 2,4,6

    c2 := Circle{x: 0, y: 0, r: 5}
    c3 := Circle{0,0,5}

    c3.x = 10
    c3.y = 15

    fmt.Println(*c1, circleArea(*c1))
    fmt.Println(c2, circleArea(c2))
    fmt.Println(c3, circleArea(c3))
}
