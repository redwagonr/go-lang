package main

import "fmt"

const boilingF float64 = 212.0

func main() {
	fmt.Println("July 24th 2019")
	fmt.Printf("Temperature in F is %gF and temperature in C is %gC", boilingF, ((boilingF - 32) * 5 / 9))
}
