package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Beginning panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	panic("Boom 💣")
}
