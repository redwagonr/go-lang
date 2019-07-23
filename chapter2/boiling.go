package main

import "fmt"

const boilingF float64 = 212.0

func main() {
	var centigrade = (boilingF - 32) * 5 / 9
	fmt.Printf("Temperature is %gF and temperature in centigrade is %gC", boilingF, centigrade)
}
