package main

import "fmt"

func main() {

	fmt.Println(mySqrt(121))

	// nums := []int{1, 2, 3, 4, 5}

	// // sum := 0

	// for c, d := range nums {
	// 	// fmt.Println(c, ":", d)
	// 	if c < d {
	// 		fmt.Println(d)
	// 	}
	// }

}

func mySqrt(x int) int {

	var a int

	for a*a < x {
		a++
	}

	for a*a > x {
		a--
	}

	return a
}
