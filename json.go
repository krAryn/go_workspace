package main

import "fmt"
import "encoding/json"
import "strings"
//import "io"

type Student struct {
	Name string	`json:"naam"`
	Class string	`json:"kaksha"`
	Roll string	`json:"kramaank"`
	Marks float32	`json:"aank"`
}

var students []Student

func main() {
	students = []Student {
		Student {Name: "Ananya", Class: "9", Roll: "23", Marks: 92.3},
		Student {Name: "Rohit", Class: "9", Roll: "24", Marks: 88.5},
		Student {Name: "Priya", Class: "9", Roll: "25", Marks: 94.7},
		Student {Name: "Karan", Class: "9", Roll: "26", Marks: 91.2},
		Student {Name: "Sneha", Class: "9", Roll: "27", Marks: 95.4},
		Student {Name: "Rahul", Class: "9", Roll: "28", Marks: 90.1},
		Student {Name: "Sara", Class: "9", Roll: "29", Marks: 93.8},
		Student {Name: "Vikas", Class: "9", Roll: "30", Marks: 87.6},
		Student {Name: "Meera", Class: "9", Roll: "31", Marks: 96.1},
		Student {Name: "Ravi", Class: "9", Roll: "32", Marks: 89.3},
	}

	var writer strings.Builder
	
	encoder := json.NewEncoder(&writer)
	encoder.Encode(students)

	fmt.Println(writer.String())
}
