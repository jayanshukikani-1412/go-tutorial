package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://www.mytruckboss.com/privacy-policy?coursename=golang&teacher=jayanshu"

func main() {
	fmt.Println("Handling URLs")
	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["teacher"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.mytruckboss.com",
		Path: "/privacy-policy",
		RawQuery: "coursename=golang&teacher=jayanshu",
	}
	fmt.Println(partsOfUrl.Scheme)
	fmt.Println(partsOfUrl.Host)
	fmt.Println(partsOfUrl.Path)
	fmt.Println(partsOfUrl.RawQuery) 
} 