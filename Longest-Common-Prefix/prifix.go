package main

import "fmt"

func main() {

	var z = []string{"shubham", "shubha", "Prajashubha"}

	// fmt.Println(Prifix(z))

	// i := 0
	var a, b []string
	c := len(z[0])

	if len(z[0]) != len(z[1]) {
		
		fmt.Println("length of c", c)


	}

	fmt.Println(z[0])
	fmt.Println(z[1])

}

func Prifix(a []string) string {

	var z []string

	for i := 0; i < len(z); i++ {
		if z[0][i] != z[0][i+1] {
			fmt.Println(z[0][i])
			fmt.Println(z[0][i+1])
		}
	}

	return "shubham"
}
