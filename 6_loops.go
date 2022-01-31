package main

import "fmt"

func main() {

	x := 0
	for x < 5 {
		fmt.Print(x, " ")
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
	names := []string{"anu", "kul", "pan", "dey"}
	for i := 0; i < len(names); i++ {
		fmt.Print(names[i])
	}

	fmt.Println()
	for _, val := range names {
		fmt.Print(val)
	}

	// Changing val in the loop wont change the real value because it is passed by value not by reference lol
}
