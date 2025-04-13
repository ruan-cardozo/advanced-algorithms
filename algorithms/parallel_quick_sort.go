package algorithms

import (
    "sync"
)

type ParallelQuickSort struct{
	
}

func (pq ParallelQuickSort) Name() string {
    return "Parallel Quick Sort"
}

func (pq ParallelQuickSort) Sort(arr []int) ([]int, int, int) {
    return pq.SortParallel(arr)
}

func (pq ParallelQuickSort) SortParallel(arr []int) ([]int, int, int) {
    var wg sync.WaitGroup
    comparisons := 0
    swaps := 0

    pq.quickSortParallel(arr, 0, len(arr)-1, &wg, &comparisons, &swaps)
    wg.Wait()

    return arr, comparisons, swaps
}

func (pq ParallelQuickSort) quickSortParallel(arr []int, low, high int, wg *sync.WaitGroup, comparisons *int, swaps *int) {
    if low < high {
        // Particiona o array e obtém o índice do pivô
        pivot, localComparisons, localSwaps := pq.partition(arr, low, high)
        *comparisons += localComparisons
        *swaps += localSwaps

        // Verifica se o tamanho do subarray é maior que o threshold
        if high-low > THRESHOLD {
            // Adiciona duas goroutines ao WaitGroup
            wg.Add(2)

            // Ordena o lado esquerdo em uma goroutine
            go func() {
                defer wg.Done()
                pq.quickSortParallel(arr, low, pivot-1, wg, comparisons, swaps)
            }()

            // Ordena o lado direito em outra goroutine
            go func() {
                defer wg.Done()
                pq.quickSortParallel(arr, pivot+1, high, wg, comparisons, swaps)
            }()
        } else {
            // Ordena sequencialmente se o tamanho for menor ou igual ao threshold
            pq.quickSortSequential(arr, low, pivot-1, comparisons, swaps)
            pq.quickSortSequential(arr, pivot+1, high, comparisons, swaps)
        }
    }
}

func (pq ParallelQuickSort) quickSortSequential(arr []int, low, high int, comparisons *int, swaps *int) {
    if low < high {
        pivot, localComparisons, localSwaps := pq.partition(arr, low, high)
        *comparisons += localComparisons
        *swaps += localSwaps

        pq.quickSortSequential(arr, low, pivot-1, comparisons, swaps)
        pq.quickSortSequential(arr, pivot+1, high, comparisons, swaps)
    }
}

func (pq ParallelQuickSort) partition(arr []int, low, high int) (int, int, int) {
    pivot := arr[high]
    i := low - 1
    comparisons := 0
    swaps := 0

    for j := low; j < high; j++ {
        comparisons++
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
            swaps++
        }
    }

    arr[i+1], arr[high] = arr[high], arr[i+1]
    swaps++

    return i + 1, comparisons, swaps
}