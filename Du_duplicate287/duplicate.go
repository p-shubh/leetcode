package main

import (
	"fmt"
	"sort"
)

func main() {

	var a []int = []int{3, 1, 3, 4, 2}

	fmt.Println("findDuplicate(a)", findDuplicate(a))

	// fmt.Println(FindDuplicate(a))

	// fmt.Println(binarySearch(a))

}

func findDuplicate(nums []int) int {
    slow, fast := 0,0
    
    for {
        slow = nums[slow]
        fast = nums[nums[fast]]
        
        if slow == fast {
            break
        }
    }
    
    slow2 := 0
    
    for {
        slow = nums[slow]
        slow2 = nums[slow2]
        
        if slow == slow2{
            break
        }
    }
    return slow2
}

func findDuplicates(nums []int) int {

	result := 0

	sort.Ints(nums)

	fmt.Println(nums)

	// occurred := make(map[int]bool, 0)
	// fmt.Println("occurred :", occurred)

	// for e := range nums {

	// 	if !occurred[e] {
	// 		occurred[nums[e]] = true

	// 		// append to result to the slice
	// 		result = append(result, nums[e])

	// 	}
	// }
	var count, i int
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			// c = nums[i]
			count++

		}
		if count > 0 {

			result = nums[i]
			count = 0
			// } else {

			// 	for z := 0; z < len(result); z++ {
			// 		if nums[i] != result[z] {
			// 			result = append(result, nums[i])
			// 			// if i == 1 {
			// 			// 	break
			// 			// }
			// 			count = 0

			// 		} else if nums[i] == result[z] {
			// 			result[z] = nums[i]
			// 			count = 0

			// 		}

			// 	}

			// }

		}

	}

	return result

}

// func FindDuplicate(nums []int) int {

// 	var l int

// 	sort.Ints(nums)

// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if i != j && nums[i] == nums[j] {
// 				l = nums[i]
// 				break
// 			}
// 		}
// 	}
// 	return l
// }

// func findDuplicate(nums []int) int {

// 	if nums == nil || len(nums) < 2 {
// 		return len(nums) - 1
// 	}

// 	start := 1
// 	end := len(nums) - 1

// 	for start < end {

// 		count := 0
// 		mid := (end-start)/2 + start

// 		for _, v := range nums {
// 			if v >= start && v <= mid {
// 				count++
// 			}
// 		}

// 		if count <= (mid-start)+1 {
// 			start = mid + 1
// 		} else {
// 			end = mid
// 		}
// 	}

// 	return end
// }

// func binarySearch(nums []int) (bool, int) {

// 	statement := false

// 	target := 8
// 	var a int
// 	Rng := len(nums)

// 	fmt.Println(Rng)

// 	sort.Ints(nums)
// 	fmt.Println(nums)
// 	// [0 2 3 4 5 6 8]

// 	low := 0

// 	high := len(nums) - 1

// 	mid := (low + high) / 2

// 	for !statement {
// 		if nums[low] == target {

// 			statement = true
// 			a = low

// 		} else if nums[high] == target {
// 			statement = true
// 			a = high
// 		} else if nums[mid] == target {
// 			statement = true
// 			a = mid
// 		} else if nums[mid] != target {
// 			low++
// 			high--
// 		}
// 	}

// 	// if nums[mid] == target {
// 	// 	statement = true
// 	// 	a = mid

// 	// for a != target {

// 	// }

// 	return statement, a
// }

// func findDuplicate(nums []int) int {

// 	// var a []int = []int{1, 3, 4, 2, 2}

// 	var l int

// 	sort.Ints(nums)

// 	low := 0

// 	high := len(nums) - 1

// 	mid := (high + low) / 2

// 	for l != 0 {

// 		if nums[low] == nums[high] {
// 			l = nums[low]
// 			return l
// 		} else if l == -1 {

// 			if nums[mid] == nums[mid-1] || nums[mid] == nums[mid+1] {
// 				l = nums[mid]
// 				return l
// 			} else {
// 				mid++
// 			}

// 		} else {
// 			low++
// 			high--
// 		}

// 	}

// 	return l
// }
