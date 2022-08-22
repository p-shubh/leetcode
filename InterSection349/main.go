/*
	Given two integer arrays nums1 and nums2, return an array of their intersection.
	Each element in the result must be unique and you may return the result in any order.

	Example 1:

	Input: nums1 = [1,2,2,1], nums2 = [2,2]
	Output: [2]

	Example 2:

	Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	Output: [9,4]
	Explanation: [4,9] is also accepted.


*/

package main

import (
	"fmt"
)

func main() {

	/*

							[1,2,2,1]
						[2,2]

									[4,9,5]
								[9,4,9,8,4]

								[1]
				[1,1]

				[2,1]
		[1,2]
	*/

	// nums1 := []int{2, 1}

	// nums2 := []int{1, 2}

	// nums1 := []int{1, 2, 2, 1}

	// nums2 := []int{2, 2}

	// nums1 := []int{4, 9, 5}

	// nums2 := []int{9, 4, 9, 8, 4}

	nums1 := []int{54, 93, 21, 73, 84, 60, 18, 62, 59, 89, 83, 89, 25, 39, 41, 55, 78, 27, 65, 82, 94, 61, 12, 38, 76, 5, 35, 6, 51, 48, 61, 0, 47, 60, 84, 9, 13, 28, 38, 21, 55, 37, 4, 67, 64, 86, 45, 33, 41}
	nums2 := []int{17, 17, 87, 98, 18, 53, 2, 69, 74, 73, 20, 85, 59, 89, 84, 91, 84, 34, 44, 48, 20, 42, 68, 84, 8, 54, 66, 62, 69, 52, 67, 27, 87, 49, 92, 14, 92, 53, 22, 90, 60, 14, 8, 71, 0, 61, 94, 1, 22, 84, 10, 55, 55, 60, 98, 76, 27, 35, 84, 28, 4, 2, 9, 44, 86, 12, 17, 89, 35, 68, 17, 41, 21, 65, 59, 86, 42, 53, 0, 33, 80, 20}

	fmt.Println(intersection(nums1, nums2))
}

// func intersection(nums1 []int, nums2 []int) []int {

// 	if len(nums1) > len(nums2) {

// 		fmt.Println("Before swapping :", "nums1 =", nums1, "nums2 =", nums2)
// 		a := nums1
// 		nums1 = nums2
// 		nums2 = a
// 		fmt.Println("After swapping :", "nums1 =", nums1, "nums2 =", nums2)

// 	}

// 	var addSlice []int
// 	// var count, i, j int

// 	// for i = 0; i <= len(nums1)-1; i++ {

// 	// 	fmt.Println("i :", i)

// 	// 	fmt.Println("len(nums1):", len(nums2)-1)

// 	// 	for j = 0; j <= len(nums2)-1; j++ {

// 	// 		fmt.Println("j :", j)
// 	// 		fmt.Println("i :", i, "nums1 :", nums1[i])
// 	// 		fmt.Println("j :", j, "nums2 :", nums2[j])

// 	// 		if len(nums1) == 1 || len(nums2) == 1 {

// 	// 			if len(addSlice) == 0 {

// 	// 				addSlice = append(addSlice, nums2[j])

// 	// 				fmt.Println("addSlice :", addSlice)

// 	// 				count = 0
// 	// 			}
// 	// 		} else if nums1[i] == nums2[j] {
// 	// 			count++
// 	// 		}

// 	// 		/*
// 	// 			nums1 = [2 2] nums2 = [1 2 2 1]
// 	// 		*/
// 	// 		fmt.Println("count :", count)

// 	// 		fmt.Println("nums1[i] :", nums1[i])

// 	// 		if count == 2 {

// 	// 			if len(addSlice) == 0 {
// 	// 				addSlice = append(addSlice, nums2[j])
// 	// 				fmt.Println("addSlice :", addSlice)
// 	// 				count = 0
// 	// 			}
// 	// 			// if nums1[i] != addSlice[k] {

// 	// 			// }
// 	// 		} else {

// 	// 			if count < 2 {
// 	// 				for _, x := range addSlice {
// 	// 					if x != nums1[i] && len(addSlice) != 0 {
// 	// 						addSlice = append(addSlice, nums2[j])
// 	// 						fmt.Println("addSlice :", addSlice)
// 	// 						count = 0
// 	// 					}

// 	// 				}

// 	// 			}

// 	// 		}

// 	// 	}

// 	// }

// 	fmt.Println("addSlice :", addSlice)

// 	return addSlice

// }

func intersection(nums1 []int, nums2 []int) []int {
	// count := 0

	if len(nums1) < len(nums2) {

		// fmt.Println("Before swapping :", "nums1 =", nums1, "nums2 =", nums2)
		a := nums1
		nums1 = nums2
		nums2 = a
		// fmt.Println("After swapping :", "nums1 =", nums1, "nums2 =", nums2)

	}

	var addSlice, addSlice2 []int
	for i := 0; i < len(nums2); i++ {
		for j := 0; j < len(nums1); j++ {

			// // fmt.Println("j :", j)
			// fmt.Println("i :", i, "nums1 :", nums1[j])
			// fmt.Println("j :", j, "nums2 :", nums2[i])

			if len(addSlice) == 0 && nums1[j] == nums2[i] {
				addSlice = append(addSlice, nums1[j])
				// fmt.Println("addSlice :", addSlice)

			} else {

				for k := 0; k < len(addSlice); k++ {

					// fmt.Println("k :", k, "addSlice[k] :", addSlice[k])
					// fmt.Println("nums1 :", nums1[j])

					if nums1[j] == nums2[i] && addSlice[k] != nums1[j] {

						addSlice = append(addSlice, nums1[j])
						// fmt.Println("addSlice :", addSlice)

					}

				}

			}

		}
	}

	inResult := make(map[int]bool)
	for _, t := range addSlice {
		if _, ok := inResult[t]; !ok {
			inResult[t] = true
			addSlice2 = append(addSlice2, t)
		}
	}
	return addSlice2

}

// for	c,d := range nums1{
// 		for e,f := range nums2{

// 		}
// 	}

// }
