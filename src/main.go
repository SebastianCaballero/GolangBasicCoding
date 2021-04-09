package main

import "fmt"

func main() {
	//Hello World
	fmt.Println("Hello World")

	//Constant Declaration
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//Variable Declaration and Area Exercise
	base := 12
	var height int = 14
	var area int = base * height

	fmt.Println("base: ", base, "; height: ", height, "; area: ", area)

	//Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("Zero Values: ", " int->", a, " float->", b, " string->", c, " bool->", d)

	//ARITHMETIC OPERATORS

	x := 50
	y := 20

	//Sum
	Sum := x + y
	fmt.Println("Sum: ", Sum)
	//Difference
	Diff := x - y
	fmt.Println("Diff: ", Diff)
	//Multiplication
	Mult := x * y
	fmt.Println("Mult: ", Mult)
	//Division
	Div := x / y
	fmt.Println("Div: ", Div)
	//Module
	Mod := x % y
	fmt.Println("Mod: ", Mod)
	//Increment
	x++
	fmt.Println("Increment: ", x)
	//Decrement
	x--
	fmt.Println("Decrement: ", x)

	//Area Exercise
	base = 24
	height = 12
	fmt.Println("Rectangule area:", base*height)

	upperBase := 20
	fmt.Println("Trapeze Area: ", ((upperBase+base)*height)/2)

	var radius float64 = 5
	fmt.Println("Circle Area: ", pi2*radius)

	// Data Types
	// var shortInt int8 = 3
	// var longInt int64 = 2313212113234256876
	// var shortFloat float32 = 3.1
	// var longFloat float64 = 3.153465212456432145668723312
	// var text string = "string"
	// var boolean bool = true
	// var complex complex128 = 10 + 8i

	//fmt Library
	fmt.Printf("Base is %d and height is %d\n", base, height)
	message := fmt.Sprintf("Hello World")
	fmt.Println(message)
}
