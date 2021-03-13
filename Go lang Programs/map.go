package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "hydrogen"
	elements["He"] = "helium"
	elements["Li"] = "lithium"
	elements["Be"] = "beryllium"
	elements["B"] = "boron"
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}
