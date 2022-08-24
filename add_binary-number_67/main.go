package main

import "fmt"

func main() {

	addBinary("11", "10")
	P := reverse("11")
	P1 := reverse("10")

	fmt.Println("P :", P, "P1 :", P1)

}

func addBinary(a string, b string) string {
	var result string

	return result

}

func reverse(s string) (t string) {
	for _, r := range s {
		t = string(r) + t
	}
	return
}
