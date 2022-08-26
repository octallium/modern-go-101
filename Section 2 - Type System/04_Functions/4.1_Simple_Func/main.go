package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	name := "Anil"

	// Immutable data types such as string, int, float, boolean
	// are passed by value.
	greet(name)
}
