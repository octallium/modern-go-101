package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		task(ch)
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("All task done!")

}

func task(ch chan<- string) { // write only
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Task: %d", i+1)
	}
}
