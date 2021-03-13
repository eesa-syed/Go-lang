package main

import "fmt"

func main() {
	var i int
	fmt.Print("enter i value:")
	fmt.Scanf("%d", &i)
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("try again")
	}
}
