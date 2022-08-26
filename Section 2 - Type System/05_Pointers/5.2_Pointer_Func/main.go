package main

import (
	"fmt"
	"strings"
)

func changeCaseByRef(s *string) {
	*s = strings.Title(*s)
}

func changeCaseByValue(s string) {
	s = strings.Title(s)
}

func greet() *string {
	s := "Hello Gopher..."
	return &s
}

func main() {
	s := "the big blue sky"
	changeCaseByValue(s)
	fmt.Println("With simple:", s)

	changeCaseByRef(&s)
	fmt.Println("With Ref:", s)

	g := greet()
	fmt.Println("Greeting:", *g)
}
