package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	City  string
	Email string
}

func (p Person) Introduce() string {
	return fmt.Sprintf("저는 %s이고 %d살입니다. 사는 곳은 %s입니다.\n이메일은 %s입니다.\n", p.Name, p.Age, p.City, p.Email)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	person := Person{
		Name:  "Alice",
		Age:   20,
		City:  "Seoul",
		Email: "sonjungchan@sonjungchan.com",
	}

	fmt.Println(person.Introduce())

	if person.IsAdult() {
		fmt.Println(person.Name, "is an adult.")
	}

	person.HaveBirthday()
	fmt.Println("after birthday:", person.Age)
}
