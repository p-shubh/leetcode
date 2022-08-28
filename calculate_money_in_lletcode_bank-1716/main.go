package main

import "fmt"

func main() {

	fmt.Println(totalMoney(4))
}

// func totalMoneys(n int) int {
// 	mod := n / 7
// 	div := n % 7
// 	return 21*mod + 7*mod*(mod+1)/2 + (2*mod+div+1)*div/2
// }

func totalMoney(n int) int {

	var total int
	var week int = 7
	var divisor int
	var rem int

	if n <= week {

		for i := 0; i <= n; i++ {
			total = total + i
		}

	} else {

		divisor = n / week

		rem = (n % week) + divisor

		// fmt.Println("divisor :", divisor)
		// fmt.Println("rem :", rem)

		// if divisor > 0 {

		for divisor >= 0 {

			for i := divisor + 1; i <= rem; i++ {
				total = total + i
			}

			rem = 6 + divisor //7 +2 -1
			// fmt.Println("rem :", rem)
			divisor--
			// fmt.Println("divisor :", divisor)
			// fmt.Println("total :", total)

		}

	}

	// }

	return total

}
