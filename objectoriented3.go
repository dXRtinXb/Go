package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Printf("%s is speaking\n", a.Name)
}

type Dog struct {
	Animal 
	Breed  string
}

func (d *Dog) Bark() {
	fmt.Printf("%s this dog is barking\n", d.Name) 
}

func main() {
	d1 := Dog{
		Animal: Animal{Name: "rex"},
		Breed:  "german shepherd",
	}

	d1.Speak() 
	d1.Bark()  
	fmt.Printf("%s یک %s است.\n", d1.Name, d1.Breed)
}