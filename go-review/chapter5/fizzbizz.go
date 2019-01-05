// Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)

package main

import "fmt"

func main () {

    var printval string
    const fizz = "Fizz"
    const bizz = "Bizz"

    for i := 1; i <= 20; i++ {
        printval = ""
        if i % 3 == 0 {
            printval = printval + fizz
        }
        if i % 5 == 0 {
            printval = printval + bizz
        }

        if printval == "" {
            fmt.Println(i)
        } else {
            fmt.Println(printval)
        }
    }

}
