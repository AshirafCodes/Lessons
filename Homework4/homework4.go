package main

import "fmt"

func main() {
	// Homework 1 Boolean and Trip
	// Booking Trip
	// Decision on Trip
	// Who are going on trip
	Frank := true
	Pascle := true
	Anita := false
	// If Frank, Pascle, or Anita want to go on the trip then they go on the trip
	fmt.Println("Going for trip?", Frank || Pascle || Anita)
	fmt.Println("Is Frank going on the trip?", Frank)
	fmt.Println("Is Pascle going on the trip?", Pascle)
	fmt.Println("Is Anita going on the trip?", Anita)
	fmt.Println("Are Anita AND Pascle going on the trip?", Anita && Pascle)
	fmt.Println("Are Anita AND Frank going on the trip?", Anita && Frank)
	fmt.Println("Is Anita or Frank going on the trip?", Anita || Frank)

	// Number: Integers and floats.
	// Integers = Whole numbers
	// Floats = decimal numbers
	// my math operators: + - * / % ()
	// % = remainder
	// / = divide
	// + = plus
	// - = subtraction
	// * = multiply
	// () = parenthesis

	// Leaving in space
	fmt.Println("")

	fmt.Println("4+4 =", 4+4)
	fmt.Println("20.5/5.5", 20.5/5.5)
	fmt.Println("4*4 =", 4*4)
	fmt.Println("4-4 =", 4-4)
	fmt.Println("4%4 =", 4%4)
	fmt.Println("7%3 =", 7%3)
	fmt.Println("8%6 =", 8%6)

	// Leaving in space
	fmt.Println("")

	fmt.Println("Below are Floats and decimals together")
	// Leaving in space
	fmt.Println("")
	// Floats || Decimal numbers only
	// Note
	fmt.Println("2.0+5 =", 2.0+5)
	fmt.Println("20.0/7 =", 20.0/7)
	fmt.Println("15*3.0 =", 15*3.0)
	fmt.Println("30.63-11.34 =", 30.63-11.34)
	fmt.Println("1.23+3.33 =", 1.23+3.33) // 4.56

	// Leaving in space
	fmt.Println("")

	fmt.Println("2+2*3 =", 2+2*3)

	// Homework 2 Boolean and States in USA
	// Ranking States
	// vote on States
	// List the best states?

	// Leaving in space
	fmt.Println("")

	iowa := true
	minnesota := true
	arizona := false
	fmt.Println("Worst states in America?", arizona && iowa && minnesota)
	fmt.Println("Is the mentioned among the best?", iowa || minnesota)
	fmt.Println("Arizona is best, right?", iowa && minnesota && arizona)

	fmt.Println("")

	// String,
	fmt.Println("Hey there")
	// Tried the fmt.println("Hey" + "there") but failed. So I used the one above.
	stringN := "n"
	stringO := "o"
	stringT := ""

	fmt.Println("")

	fmt.Println("StringN + stringO =", stringN+stringO)
	stringT = stringN + stringO
	fmt.Println("")
	fmt.Println("StringN =", stringN)
	fmt.Println("StringO =", stringO)
	fmt.Println("StringT =", stringT)

	fmt.Println("")
}
