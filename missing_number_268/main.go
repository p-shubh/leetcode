package main

import (
	"fmt"
)

func main() {

	a := []int{3, 0, 1}

	fmt.Println(missingNumber(a))

}

func missingNumber(nums []int) int {

	length := len(nums)
	fmt.Println("length :", length)
	var i int

	numsL := length - 1

	for i := 0; i <= length; i++ {

		mid := length / 2

		if mid > i {
			// fmt.Println("mid :", mid, "i:", i)
			for j := 0; j < mid; j++ {
				if nums[numsL] != i {
					return i
				}
			}

		} else if mid < i {
			for j := mid; j < numsL; j++ {
				if nums[numsL] != i {
					return i
				}
			}
		}

	}

	return i

}
