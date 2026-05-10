package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}


func (p Person) Greet() {
	fmt.Printf("سلام، من %s %s هستم.\n", p.FirstName, p.LastName)
}

func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

func main() {
	p1 := Person{
		FirstName: "Ali",
		LastName:  "Ahmadi",
		Age:       30,
	}

	p1.Greet() 

	p2 := &Person{ 
		FirstName: "mariam",
		LastName:  "Ahmadi",
		Age:       25,
	}

	p2.Greet()

	p2.SetAge(26) 
	fmt.Printf("%s Now %dyears old.\n", p2.FirstName, p2.Age)
}
