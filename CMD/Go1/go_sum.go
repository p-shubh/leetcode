package Go_sum

import "fmt"

func Go_Sum() {

	// var c []int

	// var tri bool = false

	var target int = 6

	var nums = []int{3, 2, 4}

	// for i := 0; i < len(nums); i++ {
	// 	for j := 0; j < len(nums); j++ {

	// 		c, tri = sum(nums[i], nums[j])

	// 		if tri == true {
	// 			c = []int{i, j}
	// 			break
	// 		}

	// 	}

	// 	if tri == true {
	// 		break
	// 	}

	// }

	// if tri == true {
	// 	fmt.Println(c)

	// }

	e := twoSum(nums, target)

	fmt.Println("twoSum(nums, target)", e)

}

// func sum(a int, b int) ([]int, bool) {

// 	var bool bool = false
// 	var result []int

// 	var target = 13

// 	if a+b == target {
// 		result = []int{a, b}

// 		bool = true

// 	}
// 	return result, bool
// }

func twoSum(num []int, target int) []int {

	var result []int

out:
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num); j++ {

			if i != j && num[i]+num[j] == target {

				result = []int{i, j}
				break out
			}
		}

	}

	return result
}
