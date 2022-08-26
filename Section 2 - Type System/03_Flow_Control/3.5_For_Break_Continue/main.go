package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Skipping 5...")
			continue
		}
		sum += i
		if i == 8 {
			fmt.Println("8!, breaking out of loop...")
			break
		}
	}
	fmt.Println("Sum: ", sum)
}
