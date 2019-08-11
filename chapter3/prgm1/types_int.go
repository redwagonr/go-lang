package main

import "fmt"

func variableCast() {
	fmt.Println("\nInside variableCast()")
	var apples int32 = 1
	var oranges int16 = 2
	fmt.Printf("%d\n", int32(apples)+int32(oranges))
	fmt.Println("Leaving variableCast()")
}

func bitwise() {
	fmt.Println("\nIn bitwise()")
	var x1, y1 int = 1, 0
	fmt.Printf("Value of bit clear is %d &^ %d : %d \n", x1, y1, x1&^y1)
	fmt.Println("More bitwisey")
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("x&y %08b\n", x&y)
	fmt.Printf("x|y %08b\n", x|y)
	fmt.Printf("x^y %08b\n", x^y)
	fmt.Printf("x&^y %08b\n", x&^y)
	fmt.Println("Leaving bitwise()")
}

func main() {
	/* demonstrating how overflows are handled*/
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	bitwise()
	variableCast()

}
