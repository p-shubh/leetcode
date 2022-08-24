package main

import (
	"fmt"
	"sort"
)

/*

[3,6,7,8,10]
5

*/

func main() {

	nums := []int{1, 3, 5, 6}

	target := 2

	// b := nums

	fmt.Println(searchInsert(nums, target))

	// fmt.Println(searchInserts(b, target))

}

func searchInsert(nums []int, target int) int {

	result := 0

	sort.Ints(nums)
	fmt.Println(nums)

	for a, b := range nums {

		if b == target {
			result = a

		} else if b < target {
			result = a + 1
		}

	}

	return result
}

// func searchInserts(nums []int, target int) int {

// 	result := 0

// 	actLength := len(nums) - 1

// 	if target == nums[actLength] {
// 		result = actLength
// 	} else if target == nums[0] {
// 		result = 0
// 	}

// 	sort.Ints(nums)

// 	if target == nums[actLength] {
// 		result = actLength
// 	} else if target == nums[0] {
// 		result = 0
// 	}

// 	length := (len(nums) - 1) / 2

// 	if target == length {
// 		result = length
// 	} else if target < nums[length] {
// 		c := target == nums[length-1]
// 		if c {

// 			return length - 1
// 		} else {
// 			length--
// 		}
// 	} else {
// 		c := target == nums[length+1]
// 		if c {
// 			return length + 1
// 		} else {
// 			length++
// 		}
// 	}

// 	return result

// }
