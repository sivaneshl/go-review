package main

import "fmt"

type Person struct {
    Name string
    // Person has a Name
}
func (p *Person) talk() {
    fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
    Person
    // Android is a Person; so Android has a Name
    Model string
}

func main() {
    a := new(Android)
    a.Name = "Pie"

    a.Person.talk()
    a.talk()
}
