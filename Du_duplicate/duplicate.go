package main

import "fmt"

func main() {

	var a []int = []int{1, 3, 4, 2, 2}

	fmt.Println("findDuplicate(a)", FindDuplicate(a))

}

func FindDuplicate(nums []int) int {

	var l int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] == nums[j] {
				l = nums[i]
				break
			}
		}
	}
	return l
}
