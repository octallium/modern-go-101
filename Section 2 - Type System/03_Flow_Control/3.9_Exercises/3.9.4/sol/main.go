package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
