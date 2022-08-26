package main

import "fmt"

func main() {
	ch := make(chan int)

	go task(ch)

	<-ch
	<-ch
	<-ch
	fmt.Println("All task done!")
}

func task(ch chan int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Completed task %d... \n", i)
		ch <- i
	}
}
