package Pri_fix

import "fmt"

func Prifix() string {

	var i int

	strs := []string{"shubham", "shubh", "shubhu"}

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	mx := 0
	for {
		for i = 1; i < len(strs); i++ {
			if mx >= len(strs[i-1]) || mx >= len(strs[i]) || strs[i-1][mx] != strs[i][mx] {

				return string(strs[0][:mx])
			}
		}
		mx++
	}

	fmt.Println(mx)
	fmt.Println(len(strs[i-1]))
	fmt.Println(len(strs[1]))
	fmt.Println(strs[i-1][mx])
	fmt.Println(strs[1][mx])

	return string(strs[0][:mx])
	// return string(strs[0][:mx])
}
