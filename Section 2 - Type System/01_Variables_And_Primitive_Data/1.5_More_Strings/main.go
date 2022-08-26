package main

import "fmt"

func main() {
	// strings utf-8 encoded -> int32
	// no concept of characters - read only slices of bytes or rune -> int32

	// strings can contain unicode characters
	status := "I love the rainbow 🌈"
	fmt.Println(status)

	// strings can be indexed
	fmt.Printf("The first character of status is: %c\n", status[0])

	// strings are just read-only slice of bytes
	alphabets := "abcd"
	fmt.Printf("alphabets as code points: %d\n", []byte(alphabets))

	// There is no concept of characters, instead it is only bytes or rune
	char := 'a'
	fmt.Printf("Char 'a': %v\n", char)
	fmt.Printf("Char 'a': %c\n", char)
	// rune is just an alias for type int32
	fmt.Printf("Char 'b': %d\n", char+1)

	rainbow := "🌈"
	// Even though it looks like 🌈  is a single character, the length is actually 4 bytes.
	fmt.Printf("The length of 🌈 is: %d bytes\n", len(rainbow))
	fmt.Printf("🌈 as slice of bytes: %v\n", []byte(rainbow))
	fmt.Printf("🌈 as slice of unicode characters: %U\n", []byte(rainbow))

	// looping over strings
	name := "अनिल कुलकर्णी"
	for index, runeValue := range name {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
