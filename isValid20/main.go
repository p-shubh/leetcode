package main

import (
	"fmt"
	"strings"
)

func main() {

	var g bool = true

	var c string = "shu(){}[]bham"

	d := strings.Split(c, "")

	fmt.Println(d)

	var e []string

	for i := 0; i < len(d)-1; i++ {

		for j := 0; j < len(d)-1; j++ {

			if d[i] == "(" && d[j+1] == ")" {
				e = append(e, d[i], d[j])
				g = true

			} else {
				g = false
			}
			if d[i] == "{" && d[j+1] == "}" {
				e = append(e, d[i], d[j])
				g = true

			} else {
				g = false
			}
			if d[i] == "[" && d[j+1] == "]" {
				e = append(e, d[i], d[j])
				g = true
			} else {
				g = false

			}

		}
	}

	// g = true

	fmt.Println(g)

	// fmt.Println(e)

}

// c := "s{}()[]P"

// var d []string

// e := strings.Split(c, "")

// for i := 0; i < len(e); i++ {
// 	for j := 0; j < len(e); j++ {
// 		if e[i] != e[j] {
// 			if e[i] == "(" || e[i] == "[" || e[i] == "{" {
// 				// fmt.Println("e", i, e[i])

// 				d = append(d, e[i])

// 				// if condition {

// 				// }

// 			}
// 		}
// 	}

// }

// fmt.Println(d)

// // d := strings.Split(c[0], "")

// // fmt.Println(d)

// // fmt.Println(c[0])

// ==========================================================================

// type Stack struct {
// 	Item []string
// }

// func (s *Stack) Push(i string) {
// 	s.Item = append(s.Item, i)

// 	fmt.Println(strings.Split(s.Item[len(s.Item)-1], ""))
// }

// func main() {

// 	myStack := Stack{}

// 	myStack.Push("shu(){}[]bham")

// }
