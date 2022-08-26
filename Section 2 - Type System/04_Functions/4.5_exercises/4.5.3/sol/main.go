package main

import "fmt"

func getSlice(nums []float64) ([]float64, bool, error) {
	isBelow10 := true
	newNums := make([]float64, len(nums))
	copy(newNums, nums)

	for i := 0; i < len(newNums); i++ {
		if newNums[i] < 0 {
			return []float64{}, isBelow10, fmt.Errorf("Validation Error: only positive numbers are allowed")
		} else if newNums[i] > 10 {
			isBelow10 = false
		}
		newNums[i] = newNums[i] * 2
	}

	return newNums, isBelow10, nil
}

func printSlice(nums []float64, b bool, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Are all below 10.0:", b)
		fmt.Println("Formatted Slice:", nums)
	}
}

func main() {
	n := []float64{1.2, 3.4, -4, 12.98, 45.0}
	printSlice(getSlice(n))
	fmt.Println("Original n:", n)

	m := []float64{1.2, 3.4, 4.67, 12.98, 23.18}
	printSlice(getSlice(m))
	fmt.Println("Original m:", m)
}
