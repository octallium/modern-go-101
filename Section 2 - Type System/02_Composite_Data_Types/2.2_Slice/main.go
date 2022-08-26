package main

import "fmt"

func main() {
	// declare a slice of type int
	var nums []int
	fmt.Println(nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))

	// add element to slice
	nums = append(nums, 2)
	fmt.Println(nums)
	fmt.Printf("Capacity: %d\tLength: %d\n", cap(nums), len(nums))

	// slice literal
	names := []bool{true, true, true, false}
	fmt.Println(names)

	// declaring a array
	cities := [5]string{"Mumbai", "Delhi", "Toronto", "Perth", "Oslo"}
	cityChoices := cities[0:3]
	fmt.Println(cityChoices)

	// Slice points to the underlying array
	// We will study more about slice in `slice internals`
	cities[1] = "Dublin"
	fmt.Println(cityChoices)
}
