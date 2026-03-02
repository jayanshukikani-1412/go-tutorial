package main

import (
	"io"
	"fmt"
	
	"net/http"
)

const url = "https://www.mytruckboss.com/privacy-policy"

func main() {
	fmt.Println("WEB REQUEST")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes) 
	fmt.Println(content)
}
