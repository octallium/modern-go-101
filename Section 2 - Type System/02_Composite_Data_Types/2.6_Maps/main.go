package main

import (
	"fmt"
)

func main() {
	// var contacts map[string]string
	// fmt.Println(contacts)
	// contacts["Anil"] = "Kulkarni"

	// Create and initialize a map
	contacts := make(map[string]string)
	// inserting values
	contacts["Anil"] = "anil@example.com"
	contacts["Teja"] = "teja@example.com"
	fmt.Println(contacts)
	fmt.Println("Length:", len(contacts))

	// Maps Literal
	contacts2 := map[int]string{
		1: "teja@example.com",
		2: "anil@example.com",
	}
	fmt.Println(contacts2)

	// Map with different data types
	contactNums := make(map[string][]int)
	contactNums["Anil"] = []int{44444, 22222}
	contactNums["Teja"] = []int{33333, 11111}
	fmt.Println(contactNums)

	// Key Types - All except Slice, Map & Function
	// Struct & Array elements should be comparable
}
