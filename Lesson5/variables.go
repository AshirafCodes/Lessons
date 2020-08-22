package main

import (
	"fmt"
)

func main() {
	// In Go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

	// 1. var declares 1 or more variables.
	var a = "initial"
	fmt.Println("a =", a)
	// You can reassign a variable with a value of the same type but you can't reassign a variable with a value of a different type.
	a = "something 2"
	fmt.Println("a =", a)

	// 2. You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	var b2, c2 int = 1, 2
	fmt.Println("b2 =", b2)
	fmt.Println("c2 =", c2)

	// 3. Go will infer the type of initialized variables.
	var d = true
	fmt.Println("d =", d)
	var d2 = "string"
	fmt.Println("d2 =", d2)

	// 4. Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	var e int
	fmt.Println("e =", e)

	// 5. The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	f := "apple"
	fmt.Println("f =", f)
}
