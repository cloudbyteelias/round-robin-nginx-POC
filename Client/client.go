package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const CONN string = "http://localhost:8080"

func main() {

	for {
		fmt.Println(SendRequest())
		time.Sleep(4 * time.Second)
	}

}

func SendRequest() string {
	response, err := http.Get(CONN)

	if err != nil {
		log.Fatal("Error get request")
		fmt.Printf("\t %v", response)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Request error status code is [%v]", response.StatusCode)
	}

	return "🔥order sent successfully🔥"

}
