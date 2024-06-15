package main

import "fmt"
import "encoding/json"
import "strings"
//import "io"

type Marks struct {
	Math int	
	Eng int		
	Phy int		
	Chem int	
	Cs int		
}

type AllSemMarks []Marks 

type Student struct {
	Name string	`json:"naam"`
	Class string	`json:"kaksha"`
	Roll string	`json:"kramaank"`
	AllSemMarks	`json:"aank"`
}

var students []Student

func main() {
	students = []Student {
		Student {Name: "Aryan", Roll: "22", Class: "9", AllSemMarks: []Marks{Marks{90,91,92,93,94}, Marks{81,82,83,84,85}}},
		Student {Name: "Ananya", Class: "9", Roll: "23"},
		Student {Name: "Rohit", Class: "9", Roll: "24",},
		Student {Name: "Priya", Class: "9", Roll: "25"},
		Student {Name: "Karan", Class: "9", Roll: "26"},
		Student {Name: "Sneha", Class: "9", Roll: "27"},
		Student {Name: "Rahul", Class: "9", Roll: "28"},
		Student {Name: "Sara", Class: "9", Roll: "29"},
		Student {Name: "Vikas", Class: "9", Roll: "30"},
		Student {Name: "Meera", Class: "9", Roll: "31"},
		Student {Name: "Ravi", Class: "9", Roll: "32"},
	}
	
	fmt.Printf("%T: %v\n**********************************************************\n\n", students, students)

	var writer strings.Builder
	
	encoder := json.NewEncoder(&writer)
	encoder.Encode(&students)

	fmt.Println(writer.String())
	
	fmt.Printf("\n**********************************************************\n\n")
	
	students_copy := []Student{}
	reader := strings.NewReader(writer.String())

	decoder := json.NewDecoder(reader)

	decoder.Decode(&students_copy)

	fmt.Printf("%T: %v\n**********************************************************\n\n", students_copy, students_copy)
}
