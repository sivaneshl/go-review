package main

import (
    "fmt"
    "math"
)

type Circle struct {
    x, y, r float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

type Shape interface {
    area() float64
}

func totalArea(shapes ...Shape) float64 {
    totalarea := 0.0
    for _, s := range(shapes) {
        totalarea += s.area()
    }
    return totalarea
}

func main() {
    c1 := Circle{0,0,5}
    c2 := Circle{0,0,7}
    r1 := Rectangle{2,4,6,8}

    fmt.Println(totalArea(&c1, &c2, &r1))
}
