package random_numbers

import (
	"math/rand"
)

func GenerateRandomNumbers(maximumNumber int) []int {

	numbers := make([]int, maximumNumber)

	for i := 0; i < maximumNumber; i++ {
		randomNumber := rand.Intn(maximumNumber)
		numbers[i] = randomNumber
	}

	return numbers
}