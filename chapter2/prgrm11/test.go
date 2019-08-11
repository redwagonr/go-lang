package main

import "fmt"

func toUpper(value string) {
	for _, x := range value {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}

}

func main() {
	toUpper("hello tuesday")

}
