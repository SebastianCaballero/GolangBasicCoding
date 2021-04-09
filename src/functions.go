package main

import (
	"fmt"
	"math"
)

func firstFunction(message string) {
	fmt.Println(message)
}

func threeArguments(first, second int, third string) {
	fmt.Printf("Two numbers: %d and %d. One string: %s\n", first, second, third)
}

func returningFunction(number int) int {
	return number * 2
}

func doubleReturn(number int) (c, d int) {
	return number, number * 2
}

func squareArea(base float64) float64 {
	return base * base
}
func rectangleArea(base, height float64) float64 {
	return base * height
}
func trapezeArea(base, upperBase, height float64) float64 {
	return ((upperBase + base) * height) / 2
}
func circleArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func main() {
	fmt.Println("Function program initialized!")

	message := "First message in a function example"
	firstFunction(message)

	threeArguments(15, 7, "String!")

	fmt.Println(returningFunction(4))

	num, twiceNum := doubleReturn(2)
	fmt.Printf("%d multiplied by 2 is: %d\n", num, twiceNum)

	_, twiceEight := doubleReturn(8)
	fmt.Printf("8 multiplied by 2 is: %d\n", twiceEight)

	base := 23.12
	height := 14.78
	upperBase := 18.32
	radius := 5.39

	square := squareArea(base)
	rectangle := rectangleArea(base, height)
	trapeze := trapezeArea(base, upperBase, height)
	circle := circleArea(radius)

	fmt.Printf("Square area of base %g: %g\n", base, square)
	fmt.Printf("Rectangle area with base %g and height %g: %g\n", base, height, rectangle)
	fmt.Printf("Trapeze area of base %g, upperBase %g and height %g: %g\n", base, upperBase, height, trapeze)
	fmt.Printf("Circle area of radius %g: %g\n", radius, circle)

}
