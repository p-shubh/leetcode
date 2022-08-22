package main

import "fmt"

func unique(s []string) []string {

	fmt.Println("s :", s)
	inResult := make(map[string]bool)
	fmt.Println("inResult :", inResult)

	var result []string
	fmt.Println("result :", result)

	for _, str := range s {
		fmt.Println("str :", str)

		if _, ok := inResult[str]; !ok {
			fmt.Println("ok :", ok)

			inResult[str] = true
			fmt.Println("inResult :", inResult)

			result = append(result, str)
			fmt.Println("result :", result)

		}
	}
	return result
}

func main() {
	fmt.Println(unique([]string{"abc", "cde", "efg", "efg", "abc", "cde"}))
}
