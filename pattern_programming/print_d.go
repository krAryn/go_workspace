package main

import "fmt"

func main() {
	var d int

	fmt.Scanln(&d)
	
	for i := 0; i < d; i++ {
		fmt.Print("*")
		if i == 0 || i == d - 1 {
			for j := 1; j < d - 1; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 1; j < d - 1; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}
