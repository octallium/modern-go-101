package main

import (
	"fmt"
	"log"
)

func zeroDivision(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovering from:", err)
		}
	}()
	return a / b
}

func main() {
	zeroDivision(3, 0)
	fmt.Println(zeroDivision(4, 2))
}
