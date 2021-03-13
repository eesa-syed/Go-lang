package main

import "fmt"

func main() {
	var x int
	fmt.Print("enter number:")
	fmt.Scanf("%d", &x)
	if x%2 == 0 {
		fmt.Println("the given number is even")

	} else {
		fmt.Println("the given number is odd")
	}
}
