package arraysslices

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func SumAll(numbers1, numbers2[]int) {
func SumAll(numbersToSum ...[]int) []int {
	fmt.Println(numbersToSum)

	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	fmt.Println(sums)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
