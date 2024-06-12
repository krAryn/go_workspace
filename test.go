package main

import "fmt"
import "os"
import "log"
import "strconv"

type Rect struct {
	length int
	breadth int
}

func (rec Rect) Perimeter() string {
	return fmt.Sprintf("The perimeter of the Rectangle is: 2(l + b) = %.3v\n", 2 * (rec.length + rec.breadth))
}

func (rec Rect) Area() string {
	return fmt.Sprintf("The area of the Rectangle is: l x b = %.3v\n", rec.length * rec.breadth)
} 

func main() {
	if (len(os.Args) != 3) {
		log.Fatal("Error!!!")	
	}

	l, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	
	obj := Rect{length: l, breadth: b,}
	fmt.Println(obj.Perimeter())
} 
