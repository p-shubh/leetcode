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

	nums := []int{3, 6, 7, 8, 10}

	target := 5

	fmt.Println(searchInsert(nums, target))

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
