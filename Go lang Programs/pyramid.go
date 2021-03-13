package main

import "fmt"

func main() {
	var x, y, n int
	fmt.Print("enter number of rows: ")
	fmt.Scanf("%d", &n)
	for x = 1; x <= n; x++ {
		for y = 1; y <= x; y++ {
			print(" * ")
		}
		fmt.Println(" ")
	}
}
