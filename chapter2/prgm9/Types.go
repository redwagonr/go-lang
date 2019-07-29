package main

import "fmt"

// Celsius ...
type Celsius float64

// Fahrenheit ...
type Fahrenheit float64

func cToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func fToc(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Println("This program uses to show the user defined types in Go")
	fmt.Println("Temperature in Fahrenheit is: ", cToF(100))
	fmt.Println("Temperature in Celsius is: ", fToc(212))
}
