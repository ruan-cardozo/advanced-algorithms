package main

import (
	"advanced-algorithms/algorithms"
	"advanced-algorithms/random_numbers"
	"advanced-algorithms/strategy"
	"advanced-algorithms/utils"
	"fmt"
	"os"
	"strconv"
)

func main() {

    if len(os.Args) < 4 {
        fmt.Println("Uso: go run main.go <quantidade_de_numeros> <quantidade_de_execução> <algoritmo específico ou todos>")
        fmt.Println("Algoritmos disponíveis: bubble_sort, bubble_sort_improved, insertion_sort, selection_sort, all")
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

    algorithmsMap := map[string]strategy.SortStrategy{
        "bubble_sort":         algorithms.BubbleSortStruct{},
        "bubble_sort_improved": algorithms.BubbleSortImproved{},
        "insertion_sort":      algorithms.InsertionSort{},
        "selection_sort":      algorithms.SelectionSort{},
    }

    if algorithm == "all" {
        for name, strategyByUser := range algorithmsMap {
            fmt.Printf("Executando %s...\n", name)
            sorter := strategy.NewSorter(strategyByUser)
            var totalDuration float64

            for i := 0; i < numExecutions; i++ {
                numberToSort := random_numbers.GenerateRandomNumbers(numElements)
                duration := sorter.ExecuteSort(utils.Clone(numberToSort))
                totalDuration += duration
                fmt.Println("--------------------------------------------------")
            }
            averageDuration := totalDuration / float64(numExecutions)
            fmt.Printf("Tempo médio de execução (ms): %.6f\n", averageDuration)
        }
        return
    }

    strategyByUser, exists := algorithmsMap[algorithm]

    if !exists {
        fmt.Println("Algoritmo não reconhecido. Algoritmos disponíveis: bubble_sort, bubble_sort_improved, insertion_sort, selection_sort, all")
        return
    }

    sorter := strategy.NewSorter(strategyByUser)
    var totalDuration float64

    for i := 0; i < numExecutions; i++ {

        numberToSort := random_numbers.GenerateRandomNumbers(numElements)

        duration := sorter.ExecuteSort(utils.Clone(numberToSort))
        totalDuration += duration
        fmt.Println("--------------------------------------------------")
    }
    averageDuration := totalDuration / float64(numExecutions)
    fmt.Printf("Tempo médio de execução (ms): %.6f\n", averageDuration)
}