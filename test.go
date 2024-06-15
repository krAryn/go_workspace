package main

import (
	"fmt"
	"bytes"
	"encoding/json"
//	"io"
)

type Person struct {
	Name string
	Age int
}

var data []byte = []byte(`[{"Name": "Emily", "Age": 27},
			{"Name": "John", "Age": 34},
			{"Name": "Sophia", "Age": 29},
			{"Name": "Michael", "Age": 40},
			{"Name": "Emma", "Age": 25},
			{"Name": "Daniel", "Age": 31},
			{"Name": "Olivia", "Age": 23},
			{"Name": "Liam", "Age": 28},
			{"Name": "Ava", "Age": 30},
			{"Name": "James", "Age": 35}]`)

var json_data bytes.Buffer

func main() {
	json_data.Write(data)
	//fmt.Println(json_data.String())
	
	decoder := json.NewDecoder(&json_data)
	decoded_data_slice := []Person{}

//	var person Person

//	for {	
//		err := decoder.Decode(&person)
//		if err == io.EOF {
//			break;
//		} else {
//			fmt.Println(err)
//			decoded_data_slice = append(decoded_data_slice, person)
//		}
//	}
	

	decoder.Decode(&decoded_data_slice)

	fmt.Println("decode slice: ", decoded_data_slice)
}
