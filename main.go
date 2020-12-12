package main

import (
	"fmt"
)

func GetText() string {
	return "Hello world"
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println(GetText())
}
