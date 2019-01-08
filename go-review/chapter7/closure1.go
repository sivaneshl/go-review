package main

import "fmt"

func evengenerator() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

func main() {
	nextEven := evengenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
