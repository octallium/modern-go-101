package main

import "fmt"

func main() {
	num := 6
	if num > 10 && num%2 == 0 {
		fmt.Println("num is divisible by 2")
	} else if num > 10 || num%3 == 0 {
		fmt.Println("num is less than 10 and divisible by 3")
	}
}
