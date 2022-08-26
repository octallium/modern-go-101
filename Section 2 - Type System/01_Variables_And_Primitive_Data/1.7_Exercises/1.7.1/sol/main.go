package main

import "fmt"

func main() {
	name := "Anil Kulkarni"
	city := "Toronto"
	var pincode int = 100077
	time := 10.20

	fmt.Printf("Hi, my name is %s and I stay in %s (%d). Its currently %.2f a.m here\n", name, city, pincode, time)
}
