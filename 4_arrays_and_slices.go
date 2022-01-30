package main

import "fmt"

func main() {
	// Arrays
	var ages [3]int = [3]int{20, 25, 30}
	agesTwo := [4]string{"anu", "kul", "pan", "dey"}

	fmt.Println(ages, len(ages))
	fmt.Println(agesTwo)

	// Slices
	var scores = []int{1, 3, 4, 5}
	scores = append(scores, 3)

	fmt.Println(scores)
}
