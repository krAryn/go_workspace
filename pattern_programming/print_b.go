package main

import "fmt"

func main() {
	var b int
	
	fmt.Scanln(&b)

	for i := 0; i < b; i++ {
		fmt.Print("*")
		if i == 0 || i == b/2 || i == b - 1 {
			for j := 1; j < b - 1; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 1; j < b - 1; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}
