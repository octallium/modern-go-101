package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	now := time.Now()

	go func() {
		defer close(ch)
		task(ch)
	}()
	msg := <-ch
	fmt.Println(msg)
	fmt.Println("elasped:", time.Since(now))
}

func task(ch chan string) {
	ch <- "Task 1..."
}
