package strategy

import (
	"fmt"
	"time"
)

type SortingStrategy interface {
	Sort([]int) ([]int, int, int)
}

type Sorter struct {
	Strategy SortingStrategy
}

func (s *Sorter) SetStrategy(strategy SortingStrategy) {
	s.Strategy = strategy
}

func (s *Sorter) ExecuteSort(arr []int) {
    start := time.Now()
    sortedArray, comparisons, swaps := s.Strategy.Sort(arr)
    duration := time.Since(start).Microseconds()
    // Limitar a impressão do array
    const maxPrintElements = 10
    if len(sortedArray) > maxPrintElements {
        fmt.Printf("Sorted Array: %v ... %v\n", sortedArray[:5], sortedArray[len(sortedArray)-5:])
    } else {
        fmt.Println("Sorted Array:", sortedArray)
    }
    fmt.Println("Execution Time (ms):", duration)
    fmt.Println("Comparisons:", comparisons)
    fmt.Println("Swaps:", swaps)
}