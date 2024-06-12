package main

import "fmt"
import "os"
import "log"
import "strconv"

type Rect struct {
	length int
	breadth int
}

func (rec *Rect) Operate(name string, op func() int) string {
	result := op()
	return fmt.Sprintf("The %s of the Rectangle is: %v", name, result)
} 

func main() {
	if (len(os.Args) != 3) {
		log.Fatal("Error!!!")	
	}

	l, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	
	obj := Rect{length: l, breadth: b,}
	st := obj.Operate("Area", func() int {
		return (obj.length * obj.breadth)
	})

	fmt.Println(st)
} 
