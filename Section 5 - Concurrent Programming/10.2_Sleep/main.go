package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Elapsed:", time.Since(now))
	fmt.Println("Done...")
}

func task() {
	fmt.Println("task 1")
}
