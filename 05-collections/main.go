package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Println("array:", numbers)
	fmt.Println("first number:", numbers[0])

	names := []string{"Alice", "Bob", "Carol"}
	names = append(names, "Dave")

	fmt.Println("slice:", names)
	fmt.Println("slice length:", len(names))

	scores := map[string]int{
		"Alice": 90,
		"Bob":   75,
	}
	scores["Carol"] = 88
	scores["손정찬"] = 50

	aliceScore, ok := scores["Alice"]
	if ok {
		fmt.Println("Alice score:", aliceScore)
	}

	sum := 0
	for name, score := range scores {

		fmt.Printf("%s: %d점 더하기 \n ", name, score)
		sum += score
	}

	println("평균 점수 : ", sum/len(scores))
}
