package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := Person{
		FirstName: "علی",
		LastName:  "رضایی",
		Age:       30,
	}

	fmt.Println("نام:", p1.FirstName)
	fmt.Println("سن:", p1.Age)
}
