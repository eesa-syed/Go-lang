package main

import "fmt"

func main() {
	var x []int = 50
	var n int
	var i int
	fmt.Print("enter number of elements: ")
	fmt.Scanf("%d", &n)
	fmt.Println("enter elelments in array:")
	for i = 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	var lar int
	for i = 0; i < n; i++ {
		if x[i] < x[i+1] {
			lar = x[i+1]
		}
	}
	for i = 0; i < n; i++ {
		fmt.Println("the largest element is:%d", lar)
	}
}
