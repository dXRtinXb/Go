package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := Person{
		FirstName: "Ali",
		LastName:  "Asghar",
		Age:       159,
	}

	fmt.Println("نام:", p1.FirstName)
	fmt.Println("سن:", p1.Age)
}
