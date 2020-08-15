package main

import "fmt"

func output(path string) string {
	return "This path is " + path
}

func main() {

	// Go has various value types including strings, integers, floats, booleans, etc. Here are a few basic examples.

	// Strings, which can be added together with +.
	fmt.Println("hello " + "world")
	stringA := "a"
	stringB := "b"
	stringC := ""
	fmt.Println("stringA + stringB =", stringA+stringB)
	stringC = stringA + stringB
	fmt.Println("stringA =", stringA)
	fmt.Println("stringB =", stringB)
	fmt.Println("stringC =", stringC)
	fmt.Println(output("/"))
	fmt.Println(output("/home"))

	// Number: Integers and floats.
	// Integers = whole numbers
	// Floats = decimal numbers
	// Math Operators: + - * / % ()
	//  % = remainder
	//  / = divide
	//  + = plus
	//  * = multiply
	// () = parenthesis
	fmt.Println("2+2 =", 2+2)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("negatives", -1)
	fmt.Println("2*2 =", 2*2)
	fmt.Println("2-2 =", 2-2)
	fmt.Println("2%2 =", 2%2)
	fmt.Println("3%2 =", 3%2)

	// Booleans, with boolean operators as you’d expect.
	// && = and
	// || = or
	//  ! = not
	fmt.Println("false =", true && false)
	fmt.Println("true  =", true && true)
	fmt.Println("false =", false && false)
	fmt.Println("true  =", true || false)
	fmt.Println("true  =", true || true)
	fmt.Println("false =", false || false)
	fmt.Println("false =", !true)
	fmt.Println("true  =", !false)

	// Example 1 Booleans and Tacos
	// order dinner
	// vote on dinner
	// Do you want tacos?
	bob := false
	sally := false
	steve := false
	fmt.Println("Getting Tacos?", bob || sally || steve)

}
