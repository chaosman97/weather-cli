package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	city := os.Args[1]
	res, err := http.Get("http://wttr.in/" + city + "?format=3")
	if err != nil {
		fmt.Printf("not available wttr.in %v", err)
		panic(-1)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		panic(-1)
	}

	// Print the response body as a string
	fmt.Println(string(body))
}
