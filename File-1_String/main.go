package main

import "fmt"

func main() {

	fmt.Println("This is a string")

	for i := 0; i < 4; i++ {
		fmt.Println("i is", i)
	}
	rooms := []string{"Will", "Win", "Doug", "Bob"}
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room == "Win" {
			fmt.Println(i, room)
		}
	}

}
