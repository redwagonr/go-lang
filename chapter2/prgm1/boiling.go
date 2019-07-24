package main

import "fmt"

const boilingF float64 = 212.0

func ftoc(temperature float64) float64 {
	return (temperature - 32) * 5 / 9
}

func main() {
	fmt.Printf("Temperature is %gF and temperature in centigrade is %gC", boilingF, ftoc(boilingF))
}
