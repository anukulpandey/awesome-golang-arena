package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hey!"
	fmt.Println(strings.Contains(greeting, "y"))

	fmt.Println(strings.ReplaceAll(greeting, "y", "llo"))

	ages := []int{4, 1, 3, 2, 5}
	sort.Ints(ages)
	fmt.Println(ages)

}
