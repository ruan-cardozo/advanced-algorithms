package strategy

import (
	"fmt"
	"time"
)

type SortStrategy interface {
	Sort([]int) ([]int, int, int)
}

type Sorter struct {
	Strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
    return &Sorter{Strategy: strategy}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.Strategy = strategy
}

func (s *Sorter) ExecuteSort(arr []int) float64 {
	start := time.Now()
	sortedArray, comparisons, swaps := s.Strategy.Sort(arr)
	duration := time.Since(start).Seconds() * 1000

	// Limitar a impressão do array
	const maxPrintElements = 10
	if len(sortedArray) > maxPrintElements {
		fmt.Printf("Sorted Array: %v ... %v\n", sortedArray[:5], sortedArray[len(sortedArray)-5:])
	} else {
		fmt.Println("Sorted Array:", sortedArray)
	}

    fmt.Printf("Execution Time (ms): %.6f\n", duration)
	fmt.Println("Comparisons:", comparisons)
	fmt.Println("Swaps:", swaps)

	return duration
}
