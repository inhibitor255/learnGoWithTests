package main

func Sum(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] += Sum(numbers)
	}

	return sums
}
