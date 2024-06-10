package main

import "fmt"

func main() {
	var g int

	fmt.Scanln(&g)
	
	for i := 0; i < g; i++ {
		fmt.Print("*")
		if i == 0 || i == g - 1 {
			for j := 1; j < g - 1; j++ {
				fmt.Print("*")
			}
		} else if i > g/2 {
			for j := 1; j < g - 1; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
		} else if i == g/2 {
			for j := 1; j < g - 1; j++ {
				if j > g/2 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
