package main

import "fmt"

func main() {
	go task()
	fmt.Println("Done without task 1...")
}

func task() {
	fmt.Println("task 1")
}
