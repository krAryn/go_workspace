package main

import "fmt"
import "net/http"
import "time"
import "context"
import "log"
import "io"

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://www.google.com", nil)
	
	if err != nil {
		log.Fatal("Error in creating http request!!\n", err)
	}
	
	res, err := client.Do(req)

	if err != nil {
		log.Fatal("Error in getting response!!\n", err)
	}

	data := make([]byte, 2048)
	
	for {
		n, err := res.Body.Read(data)

		fmt.Print(string(data[:n]))
	
		if err == io.EOF {
			break
		}
	}
}
