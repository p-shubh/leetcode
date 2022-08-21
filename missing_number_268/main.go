package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 2}

	// fmt.Println(missingNumber(a))

	fmt.Println("missingNumbers", missingNumbers(a))

}

func missingNumber(nums []int) int {

	length := len(nums)

	sort.Ints(nums)

	var result int

	fmt.Println("shorted :", nums)

	for c, d := range nums {

		if length == 1 {
			if d > 0 {
				result = nums[c] - 1
				return result
			} else {
				result = nums[c] + 1
				return result
			}

		} else if c <= len(nums) {

			fmt.Println("c :", c, "d:", d)
			if c != nums[c] {
				nums[c] = c
				result = nums[c]
				break

			} else if nums[length-1] != length {
				result = nums[c] + 1
			}
		}

	}

	return result

}

func missingNumbers(nums []int) int {
	var sum int
	var ind int
	for i, v := range nums {
		sum += v
		ind += i + 1
	}
	return ind - sum
}
