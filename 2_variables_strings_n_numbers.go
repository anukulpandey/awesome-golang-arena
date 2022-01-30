package main

import "fmt"

func main() {
	// strings

	var nameOne string = "anukul"
	fmt.Println(nameOne)

	var nameTwo = "string2"
	fmt.Println(nameTwo)

	fmt.Println(nameOne, nameTwo)

	var nameThree string
	nameThree = "ok [!]"

	nameFour := "anu" //can be defined like this in any function only

	fmt.Println(nameThree, nameFour)

	//Integers

	var ageOne int = 69
	var ageTwo = 74
	ageThree := 44

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 14 //signed int , range -128 to 127

	fmt.Println(numOne)

	var numTwo uint8 = 129 //unsigned int, range 0 to 255

	fmt.Println(numTwo)

	// float numbers

	var floatingNum float32 = 232.22 //but we will be using float64 for most of the times because it is more precise

	fmt.Println(floatingNum)
}
