package main

import "fmt"

/*
[3,3,3]
3

*/

func main() {

	nums := []int{3, 3, 3, 3}

	target := 3

	fmt.Println("search range", searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {

	var result []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = append(result, i)
		}

		if len(result) == 1 {
			result = append(result, result[0])
		}

		if len(result) > 2 {
			result = []int{result[0], result[1] + 1}
		}
	}

	if len(result) == 0 {
		result = []int{-1, -1}
	}

	return result

}
