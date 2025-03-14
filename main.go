package main

import (
	"advanced-algorithms/algorithms"
	"advanced-algorithms/random_numbers"
	"advanced-algorithms/strategy"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <quantidade_de_numeros>")
		return
	}

	numElements, err := strconv.Atoi(os.Args[1])
	if err != nil || numElements <= 0 {
		fmt.Println("Por favor, forneça um número válido maior que 0.")
		return
	}

	filename := "numbers.txt"

	err = random_numbers.GenerateRandomNumbers(numElements, filename)
	if err != nil {
		fmt.Println("Erro ao gerar números:", err)
		return
	}

	numbers, err := random_numbers.LoadNumbers(filename)
	if err != nil {
		fmt.Println("Erro ao carregar números:", err)
		return
	}

	fmt.Println("Quantidade de números gerados:", len(numbers))

	clone := func(arr []int) []int {
		newArr := make([]int, len(arr))
		copy(newArr, arr)
		return newArr
	}

	sorter := strategy.NewSorter(algorithms.BubbleSort{})
	fmt.Println("Bubble Sort:", sorter.Sort(clone(numbers)))

	sorter.SetStrategy(algorithms.BubbleSortImproved{})
	fmt.Println("Bubble Sort Melhorado:", sorter.Sort(clone(numbers)))

	sorter.SetStrategy(algorithms.InsertionSort{})
	fmt.Println("Insertion Sort:", sorter.Sort(clone(numbers)))

	sorter.SetStrategy(algorithms.SelectionSort{})
	fmt.Println("Selection Sort:", sorter.Sort(clone(numbers)))
}
