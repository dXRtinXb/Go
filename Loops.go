package main

import "fmt"

func main() {
	number := 0
	for i := 0 ; i<101 ;i++{
		number =number+i
	}
	fmt.Printf("the sum of numbers between 0 to 100 is %d" , number)
}