package main

import "fmt"

func main() {
	age := 69

	if age > 60 {
		fmt.Printf("Your age is %v", age)
	} else {
		fmt.Println("jawaan hai tu")
	}

	fmt.Println("\n", age > 600)
}
