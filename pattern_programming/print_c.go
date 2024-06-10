package main

import "fmt"

func main() {
	var c int
	
	fmt.Scanln(&c)

	for i := 0; i < c; i++ {
		fmt.Print("*")
		if i == 0 || i == c - 1 {
			for j := 1; j < c - 1; j++ {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
