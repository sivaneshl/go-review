package main
import "fmt"

func greatest(args ...int) int {
    big := 0
    for _,val := range(args) {
        if val > big {
            big = val
        }
    }
    return big
}

func main() {
    fmt.Println(greatest(55,78,22,89,665,2342,96,1,867,31,536,996))
}
