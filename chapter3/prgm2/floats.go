package main

import (
	"fmt"
	"math"
)

func exponent() {

	for i := 0; i < 8; i++ {
		fmt.Printf("x= %d e = %8.8f\n", i, math.Exp(float64(i)))
	}

}

func main() {
	fmt.Println("Happy Thursday")
	exponent()
}
