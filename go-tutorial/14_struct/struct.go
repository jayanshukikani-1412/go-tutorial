package main

import "fmt"

type person struct {
	name string
	age int
	email string
}

func main() {

	p1 := person{name: "John", age: 30, email: "john@example.com"}
	fmt.Println(p1)

	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Println(p1.email)

	p1.name = "Jane"
	fmt.Println(p1.name)

	p1.age = 31
	fmt.Println(p1.age)

	p1.email = "jane@example.com"
	fmt.Println(p1.email)
}