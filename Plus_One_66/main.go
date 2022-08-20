package main

import "fmt"

func main() {

	var digit []int = []int{8, 9}
	fmt.Println(digit)

	c := plusOne(digit)

	fmt.Println(c)

}

func plusOne(digits []int) []int {

	n := len(digits)

	var excDigits []int = []int{0}

	fmt.Println(n)

	length := n - 1

	fmt.Println("length", length)
	j := digits[length]
	fmt.Println("j", j)

	if n != 1 && ((j+1)%10) != 0 {
		digits[n-1] = j + 1
		return digits

	} else if ((j + 1) % 10) == 0 {

		for ((j + 1) % 10) == 0 {

			fmt.Println("[j+1]%10", ((j + 1) % 10))

			digits[length] = j * 0

			fmt.Println("digits[length]", digits[length])
			fmt.Println("digits", digits)

			length = length - 1
			if length == -1 {
				excDigits[0] = excDigits[0] + 1
				excDigits = append(excDigits, digits...)
				digits = excDigits
				return digits
			}

			j = digits[length]

			fmt.Println("j:", j)

			// j = digits[length]

		}

		digits[length] = digits[length] + 1

		return digits

	} else if n == 1 {

		j := digits[0]

		digits[0] = j + 1

		// if j == 0 || j == 1 {
		// 	digits[0] = j + 1
		// 	return digits
		// } else {
		// 	digits = append(digits, j*0)
		// }

	}

	return digits
}
