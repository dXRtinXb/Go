package main

import "fmt"

func main() {

	// var Sli []int
	// Sli = append(Sli, 1)
	// Sli = append(Sli, 2)

	// Sli2 := []int{1,2,3}
	// fmt.Printf("%d is the Sli's len:" , len(Sli))
	// fmt.Printf("%d is the Sli2's cap" , cap(Sli2))
	Sli := make([]string, 0)
	Sli = append(Sli , "Hello")
	Sli = append(Sli , "World")
	fmt.Println(Sli)

	Sli2 := []string{"Hello", "World"}
	fmt.Println(Sli2)
}