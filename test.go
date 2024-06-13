package main

import "fmt"

type any interface{}

func test(a interface{}) {
	another, _ := a.([]int)
	another[0] = 11
}

func main() {
	sli := []int {1,2,3,4}

	test(sli)

	fmt.Println(sli)
}
