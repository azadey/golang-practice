package main

import "fmt"

func main() {

	numbers := []int{1, 10, 20}
	sum := sumup(100, 1, 2, 3)

	anotherSum := sumup(100, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(start int, numbers ...int) int {
	sum := start

	for _, val := range numbers {
		sum += val
	}

	return sum
}
