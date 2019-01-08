package main

import "fmt"

func get_avg(values []float64) float64 {
	total := 0.0
	for _, val := range values {
		total += val
	}
	return (total / float64(len(values)))
}

func get_tot_avg(values []float64) (float64, float64) {
	total := 0.0
	for _, val := range values {
		total += val
	}
	return total, (total / float64(len(values)))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(get_avg(xs))

	fmt.Println(get_tot_avg(xs))
}
