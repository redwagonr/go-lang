package main

import "fmt"

func pointer() {

	var i int32
	fmt.Printf("Type of value i is %T\n", i)

	var p = &i
	*p++
	fmt.Printf("Type of value p which is pointer of i is %d\n", *p)
	fmt.Printf("Type of value p which is pointer of i is %d\n", p)
}

func main() {
	pointer()
}
