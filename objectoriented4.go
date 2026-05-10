package main

import "fmt"

type Speaker interface {
	Speak() string 
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s می‌گوید: هاپ هاپ!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s می‌گوید: میو میو!", d.Name)
}

func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{Name: "وفادار"}
	cat := Cat{Name: "پیشی"}

	MakeSpeak(dog) 
	MakeSpeak(cat) 
}
