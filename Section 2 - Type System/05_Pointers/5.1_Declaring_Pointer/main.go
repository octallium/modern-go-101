package main

import "fmt"

func main() {
	s := "gopher"
	fmt.Println("Address of s:", &s)

	var a *string = &s
	fmt.Printf("Value: %s\tType: %T\tPoints to address: %v\n", *a, a, a)
	fmt.Println("Address of a:", &a)

}
