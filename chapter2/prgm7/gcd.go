package main

import "fmt"

func gcd(i, j int) int {
	var div int
	div = i
	for ; div > 1; div-- {
		if i%div == 0 && j%div == 0 {
			return div
		}
	}
	return div
}

func main() {
	fmt.Print(gcd(12, 20010))
}
