package main

// Debug and clear the deadlock
import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello"
	}()

	fmt.Println(<-ch)
}
