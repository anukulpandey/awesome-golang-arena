package main

import "fmt"

func main() {
	name := "anukul"
	age := 20

	// Print
	fmt.Print("hello \n")
	fmt.Print("hello\n")

	// Println
	fmt.Println("My name is", name, "my age is", age)

	// Printf : Formatted string
	fmt.Printf("my name is %v and my age is %v", name, age)
	// %T tells the type of variable [format specifiers]
	fmt.Printf("\nage is %T\n", age)

	//Sprintf : save formatted string
	var str = fmt.Sprintf("my name is %v and my age is %v", name, age)
	fmt.Println(str)
}
