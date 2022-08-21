package main

import "fmt"

func main() {
	var digit []int = []int{}

	c := plusOne(digit)

	fmt.Println(c)

}

func plusOne(digits []int) []int {
	i := len(digits) - 1
	carry := 1
	for i >= 0 || carry > 0 {
		if i < 0 {
			digits = append([]int{carry}, digits...)
			break
		}
		sum := digits[i] + carry
		carry = sum / 10
		digits[i] = sum % 10
		i--
	}
	return digits
}
