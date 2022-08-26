package main

import (
	"fmt"
	"strings"
)

func getFormattedString(s string) (string, error) {
	if len(s) <= 5 {
		return "", fmt.Errorf("Length Validation Error: length should be 5 or more")
	}
	return strings.ToUpper(s), nil
}

func main() {
	if _, err := getFormattedString("Anil"); err != nil {
		fmt.Println(err)
	}

	if s, err := getFormattedString("watermelon"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
