package main

import "fmt"

func main() {
	var n int
	var k int = 0
	fmt.Println("enter n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		k = 0
		for space := 1; space <= n-1; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println(" ")
	}
}
