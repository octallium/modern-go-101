package main

import "fmt"

func main() {
	defer fmt.Println("Three")
	fmt.Println("One")
	fmt.Println("Two")
}
