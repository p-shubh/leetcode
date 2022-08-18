package main

import "fmt"

func main() {

	var a int = 2312

	// fmt.Println("Enter the number")

	// fmt.Scan(&a)

	result := IsPalindrome(a)

	if result {
		fmt.Println(a, "is a palindrome number")
	} else if !result {
		fmt.Println(a, " is not a palindrome number")

	}
}

func IsPalindrome(x int) bool {

	var temp int

	temp = x

	var result bool = false

	var rem int

	var res int = 0

	for temp > 0 {
		rem = temp % 10
		res = (res * 10) + rem
		temp = temp / 10
	}

	if res == x {
		result = true
	}

	return result
}
