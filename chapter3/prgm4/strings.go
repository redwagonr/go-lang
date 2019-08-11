package main

import "fmt"

func main() {

	s := "Hello World"
	fmt.Printf("Length of Hello World is: %d\n", len(s))
	fmt.Printf("Value in the last is: %d\n", s[len(s)-1])

	/* Let's take a substring */
	fmt.Println("Substring of 0 to 5 is: ", s[0:5])

	/* Substring has many flavors - Let's try a default */
	fmt.Println("Substring from 3 to end is: ", s[3:])
	fmt.Println("Substring from start to 5 is: ", s[:5])
	fmt.Println("Substring without any indexes is ", s[:])

	/* String concatentation*/
	c := "good bye " + s[5:]
	fmt.Println("Value is ", c)
}
