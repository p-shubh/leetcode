package main

import "fmt"

func main() {
	var digit []int = []int{9, 9, 9}

	c := plusOne(digit)

	fmt.Println(c)

}

func plusOne(digits []int) []int {

	n := len(digits)

	if n != 1 {

		j := digits[n-1]

		digits[n-1] = j + 1

	} else if n == 1 {

		j := digits[0]

		digits[0] = 1

		if digits[n-1]%10 == 0 {
			digits = append(digits, j*0)
			n--
		} else if j == 0 || j == 1 {
			digits[0] = j + 1
			return digits
		}

	}

	return digits
}
