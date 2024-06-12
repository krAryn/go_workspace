package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)
	
	str_cnt := n
	sp_cnt := 0

	for i := 0; i <= n/2; i++ {
			for sp := 0; sp < sp_cnt; sp++ {
				fmt.Print(" ")
			}
			for st := 0; st < str_cnt; st++ {
				fmt.Print("*")
			}
			sp_cnt++
			str_cnt = str_cnt - 2
			fmt.Println()
	}

	str_cnt = str_cnt + 4
	sp_cnt = sp_cnt - 2

	for i := n/2 + 1; i < n; i++ {
		for sp := 0; sp < sp_cnt; sp++ {
			fmt.Print(" ")
		}
		for st := 0; st < str_cnt; st++ {
			fmt.Print("*")
		}
		sp_cnt--
		str_cnt = str_cnt + 2
		fmt.Println()
	}
}
