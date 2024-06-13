package main

import "fmt"
import "strings"
import "io"

func process(reader io.Reader) {
	data := make([]byte, 2) 	// data is a byte slice of size 2
	
	for {
		count, err := reader.Read(data)
		fmt.Println("Data: ", string(data[0:count]), " Iterator: ", err)
		if (count < 2 || err == io.EOF){
			break
		}
	}
}

func main() {
	st := strings.NewReader("Hello World")
	process(st)
}
