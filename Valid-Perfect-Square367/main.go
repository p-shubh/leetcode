package main

import "fmt"

func main() {

	fmt.Println(isPerfectSquare(1))

}

func isPerfectSquare(nums int) bool {

	result := false

	if nums > 1 {
		for i := 1; i <= nums/2 && i*i <= nums; i++ {
			if i*i == nums {
				result = true
				return result
			}
		}

	} 
	// else if nums <= 1 {
	// 	for i := 1; i <= nums; i++ {
	// 		if i*i == nums {
	// 			result = true
	// 			return result
	// 		}
	// 	}
	// }

	return result

}
