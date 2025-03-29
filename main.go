package main

import (
	"advanced-algorithms/algorithms"
	tracer "advanced-algorithms/otel"
	"advanced-algorithms/random_numbers"
	"advanced-algorithms/strategy"
	"advanced-algorithms/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	shutdown := tracer.InitTracer()
	defer shutdown()

	if len(os.Args) < 4 {
		fmt.Println("Uso: go run main.go <quantidade_de_numeros> <quantidade_de_execuções> <algoritmo|algoritmos separados por virgula|all>")
		fmt.Println("Exemplo: go run main.go 100 10 bubble_sort,bubble_sort_improved")
		fmt.Println("Algoritmos disponíveis: bubble_sort, bubble_sort_improved, insertion_sort, selection_sort, merge_sort, parallel_merge_sort, quick_sort, tim_sort, heap_sort, all, ")
		return
	}

	numElements, err := strconv.Atoi(os.Args[1])
	if err != nil || numElements <= 0 {
		fmt.Println("Por favor, forneça um número válido maior que 0 para quantidade de números.")
		return
	}

	numExecutions, err := strconv.Atoi(os.Args[2])
	if err != nil || numExecutions <= 0 {
		fmt.Println("Por favor, forneça um número válido maior que 0 para quantidade de execução.")
		return
	}

	algorithm := os.Args[3]
	regex := regexp.MustCompile(`,+`)

	algorithmsMap := map[string]strategy.SortStrategy{
		"bubble_sort":          algorithms.BubbleSort{},
		"bubble_sort_improved": algorithms.BubbleSortImproved{},
		"insertion_sort":       algorithms.InsertionSort{},
		"selection_sort":       algorithms.SelectionSort{},
		"merge_sort":           algorithms.MergeSort{},
		"quick_sort":           algorithms.QuickSort{},
		"tim_sort":             algorithms.TimSort{},
		"heap_sort":            algorithms.HeapSort{},
		"parallel_merge_sort":  algorithms.ParallelMergeSort{},
	}

	if regex.MatchString(algorithm) {

		algorithmsList := strings.Split(algorithm, ",")

		for i := 0; i < numExecutions; i++ {
			numberToSort := random_numbers.GenerateRandomNumbers(numElements)

			for _, algorithm := range algorithmsList {

				strategyByUser, exists := algorithmsMap[algorithm]

				if !exists {
					fmt.Println("Algoritmos disponíveis: bubble_sort, bubble_sort_improved, insertion_sort, selection_sort, merge_sort, quick_sort, tim_sort, heap_sort, all")
					return
				}

				sorter := strategy.NewSorter(strategyByUser)
				var totalDuration float64

				duration := sorter.ExecuteSort(utils.Clone(numberToSort))
				totalDuration += duration
			}
		}
		return
	}

	if algorithm == "all" {
		for _, strategyByUser := range algorithmsMap {
			sorter := strategy.NewSorter(strategyByUser)
			var totalDuration float64

			for i := 0; i < numExecutions; i++ {
				numberToSort := random_numbers.GenerateRandomNumbers(numElements)
				duration := sorter.ExecuteSort(utils.Clone(numberToSort))
				totalDuration += duration
			}
		}
		return
	}

	strategyByUser, exists := algorithmsMap[algorithm]

	if !exists {
		fmt.Println("Algoritmo não reconhecido. Algoritmos disponíveis: bubble_sort, bubble_sort_improved, insertion_sort, selection_sort, merge_sort, quick_sort, tim_sort, heap_sort, all")
		return
	}

	sorter := strategy.NewSorter(strategyByUser)
	var totalDuration float64

	for i := 0; i < numExecutions; i++ {

		numberToSort := random_numbers.GenerateRandomNumbers(numElements)

		duration := sorter.ExecuteSort(utils.Clone(numberToSort))
		totalDuration += duration
	}
}
