package hello

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello,", name)
}

func cantSayHello(name string) {
	fmt.Println("No Hello,", name)
}
