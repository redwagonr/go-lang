package main

import "fmt"

func btoi(value bool) int {
	if value {
		return 1
	}
	return 0
}

func main() {
	fmt.Printf("Value of true is %d\n", btoi(true))
	fmt.Printf("Value of false is %d\n", btoi(false))

}
