package main

import (
	"errors"
	"fmt"
)

func greet(name string) string {
	return "안녕하세요, " + name + "님!"
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("a(10) + b(15) = ", add(10, 15))

	fmt.Println(greet("Alice"))

	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("result: %.2f\n", result)
	}

	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("result:", result)
}
