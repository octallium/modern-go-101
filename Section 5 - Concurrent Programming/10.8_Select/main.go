package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	quit := time.After(500 * time.Millisecond)

	go task(ch)

	for {
		select {
		case <-ch:
			fmt.Println(<-ch)
		case <-quit:
			fmt.Println("Time is up!")
			return
		}
	}
}

func task(ch chan<- string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- fmt.Sprintf("Task: %d", i)
	}
}
