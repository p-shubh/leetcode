package main

import "fmt"

func main() {

	var result []int

	target := 9

	var nums = []int{2, 7, 11, 15}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				result = []int{nums[i], nums[j]}
			}

		}

	}

	fmt.Println(result)
}
