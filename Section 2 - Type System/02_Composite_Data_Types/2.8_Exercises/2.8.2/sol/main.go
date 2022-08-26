package main

import "fmt"

func main() {
	emojis := make(map[string]string)
	// Make a map of 4 emoji's
	emojis["Like"] = "👍"
	emojis["Dog"] = "🐶"
	emojis["Fire"] = "🔥"
	emojis["Sparkles"] = "✨"

	// check to see if your favorite emoji is present or not
	if val, ok := emojis["Dog"]; ok == true {
		fmt.Println("Key found for: Dog", val)
	} else {
		fmt.Println("Key not found")
	}

	// delete your least favorite emoji
	delete(emojis, "Like")
	fmt.Println(emojis)

	// print out the emoji's using for-range syntax
	for key, val := range emojis {
		fmt.Printf("%s: %s\n", key, val)
	}
}
