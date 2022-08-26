package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}
