package main

import (
	"fmt"
	"leetcode/CMD/Du_duplicate"
	Go_sum "leetcode/CMD/Go1"
	Pri_fix "leetcode/CMD/Longest-Common-Prefix"
	pallindrome "leetcode/CMD/pallindrome_number"
)

func main() {

	Du_duplicate.Duplicate()

	Go_sum.Go_Sum()

	pallindrome.Palindrome()

	fmt.Println(Pri_fix.Prifix())
}
