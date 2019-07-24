package main

import "fmt"

func varInitializations() {
	// var name type = expression
	var floater float64 = 212
	//var name = expression
	var floater1 = 212.0
	// var name = complex expression
	var floater2 = (100.1 * 2.0)
	// var name
	var floater3 float64
	fmt.Printf("%T\n", floater)
	fmt.Printf("%T\n", floater1)
	fmt.Printf("%T\n", floater2)
	fmt.Printf("%T\n", floater3)
}

func multiVarInitializations() {

	//multi var same type
	var i, j int
	fmt.Printf("Type should be int, it is %T\n", i)
	fmt.Printf("Type should be int, it is %T\n", j)

	// multi var same type with initiliazations
	var i1, j1 = 2, 3
	fmt.Printf("Type should be int, it is %T\n", i1)
	fmt.Printf("Type should be int, it is  %T\n", j1)

	// multi var mixed type with initiliazations
	var i2, j2 = 2, "string"
	fmt.Printf("Type should be int, it is %T\n", i2)
	fmt.Printf("Type should be string, it is is %T\n", j2)

}

func shortVarDeclaration() {
	/* Short Variable Declarations*/

	i := 0
	fmt.Printf("Value should be int, it is %T\n", i)

	i1, j1 := 0, 1
	fmt.Printf("Value should be int, it is %T\n", i1)
	fmt.Printf("Value should be int, it is %T\n", j1)

	i2, j1 := 0, 2
	fmt.Printf("Value should be int, it is %T\n", i2)
	fmt.Printf("Value should be int, it is %T, value is %d\n", j1, j1)

}

func main() {
	fmt.Println("Different type of variables July 24th 2019")
	varInitializations()
	multiVarInitializations()
	shortVarDeclaration()

}
