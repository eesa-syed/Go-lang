package main

import "fmt"

func main() {
	xs := []float64{98, 96, 45, 65}
	fmt.Println(average(xs))
}
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
