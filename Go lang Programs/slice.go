package main

import "fmt"

func main() {
	x := []float64{1, 2, 3, 4, 5}
	y := append(x, 4, 5)
	fmt.Println(x, y)
}
