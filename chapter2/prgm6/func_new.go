package main

import "fmt"

func main() {
	//working with the new function
	var p = new(int64)
	*p = 100
	fmt.Printf("Value of *p is %d", *p)
}
