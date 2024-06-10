package main

import "fmt"

func main() {
	var a int

	fmt.Scanln(&a)

	ls := a / 2
	rs := a / 2

	for i := 0; i < a/2; i++ {
		for j := 0; j < a; j++ {
			if j == ls || j == rs || (i == a/4 && j > ls && j < rs) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		ls = ls - 1
		rs = rs + 1
		fmt.Println()
	}
}
