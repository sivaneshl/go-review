package main

import "fmt"

func main () {

    x := make(map[string]int)

    x["one"] = 10

    fmt.Println(x["one"])

    name, ok := x["one"]
    fmt.Println(name,ok)

    name, ok = x["two"]
    fmt.Println(name,ok)

    x["two"] = 20
    fmt.Println(x)


}
