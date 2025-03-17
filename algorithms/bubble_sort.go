package algorithms

import (
	"advanced-algorithms/random_numbers"
	"advanced-algorithms/strategy"
)

type BubbleSortStruct struct{}

func (b BubbleSortStruct) Sort(arr []int) ([]int, int, int) {
    n := len(arr)
    comparisons, swaps := 0, 0
    var swapped bool

    for i := 0; i < n-1; i++ {
        swapped = false
        for j := 0; j < n-1-i; j++ {
            comparisons++
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swaps++
                swapped = true
            }
        }
        if !swapped {
            break
        }
    }
    return arr, comparisons, swaps
}

func BubbleSort(numbersAmount int) {

    numbers := random_numbers.GenerateRandomNumbers(numbersAmount)

    sorter := strategy.Sorter{Strategy: BubbleSortStruct{}}
    sorter.ExecuteSort(numbers)
}