package main

import "fmt"

func tuple() (int, int) {
	fmt.Println("Tuple assignments, multiple assignments")
	var i, j, k = 1, 2, 3
	fmt.Printf("%d%d%d\n", i, j, k)
	return i, j
}

func main() {
	fmt.Println(tuple())
	var j int = 0
	_, j = tuple()
	fmt.Printf("value of j is %d", j)
}
