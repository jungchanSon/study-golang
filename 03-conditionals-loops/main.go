package main

import "fmt"

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("grade: A")
	} else if score >= 80 {
		fmt.Println("grade: B")
	} else if score >= 70 {
		fmt.Println("grade: C")
	} else {
		fmt.Println("grade: F")
	}

	day := "sat"
	switch day {
	case "sat", "sun":
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	for i := 1; i <= 5; i++ {
		fmt.Printf("%d번째 반복\n", i)
	}

	n := 0
	for n < 3 {
		fmt.Println("n =", n)
		n++
	}

	fruits := []string{"apple", "banana", "cherry"}
	for index, value := range fruits {
		fmt.Printf("[%d] %s\n", index, value)
	}
	for i := 0; i <= 10; i++ {
		fmt.Println("0부터 10: ", i)
	}
}
