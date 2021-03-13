package main

import "fmt"

func main() {
	var x int
	fmt.Print("enter nth number: ")
	fmt.Scanf("%d", &x)
	var tot int = 0
	for n := 0; n <= x; n++ {
		tot += n
		fmt.Println(tot)
	}
}
