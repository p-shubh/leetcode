package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {

	result := false

	var sum, total int

	var store []int = []int{-1}

	remainder := n % 10

	fmt.Println(remainder)

	for n > 0 {
		total = total + ((n % 10) * (n % 10))
		n = n / 10
		if n == 0 {
			sum = total
			if sum == 1 {
				result = true
				return result
			} else {
				if binarySearch(sum, store) && sum != 1 {
					return result
				}
				n = sum
				store = append(store, sum)
				total = 0
			}
		}
	}

	return result

}

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}
