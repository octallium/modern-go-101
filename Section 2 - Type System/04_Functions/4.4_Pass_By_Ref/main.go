package main

import "fmt"

func main() {
	n := []int{2, 4, 6, 8}
	fmt.Println("n:", n)

	passByRef(n)
	fmt.Println("n:", n)

	m := []int{1, 3, 5, 7}
	fmt.Println("m:", m)
	passByRefNCopy(m)
}

func passByRef(nums []int) {
	nums[0] = -2

	n2 := nums
	n2[1] = -4
}

func passByRefNCopy(nums []int) {
	n := make([]int, len(nums))
	copy(n, nums)
	n[0] = -1
	fmt.Println("Ref & Copy:", n)
}
