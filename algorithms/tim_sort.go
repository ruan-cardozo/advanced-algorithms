package algorithms

const RUN = 32

type TimSort struct{}

func (t TimSort) Sort(arr []int) ([]int, int, int) {
    comparisons, swaps := 0, 0
    timSort(arr, &comparisons, &swaps)
    return arr, comparisons, swaps
}

func insertionSort(arr []int, left, right int, comparisons, swaps *int) {
    for i := left + 1; i <= right; i++ {
        temp := arr[i]
        j := i - 1
        for j >= left && arr[j] > temp {
            *comparisons++
            arr[j+1] = arr[j]
            *swaps++
            j--
        }
        arr[j+1] = temp
        *swaps++
    }
}

func merge(arr []int, l, m, r int, comparisons, swaps *int) {
    len1, len2 := m-l+1, r-m
    left, right := make([]int, len1), make([]int, len2)
    for i := 0; i < len1; i++ {
        left[i] = arr[l+i]
    }
    for i := 0; i < len2; i++ {
        right[i] = arr[m+1+i]
    }

    i, j, k := 0, 0, l
    for i < len1 && j < len2 {
        *comparisons++
        if left[i] <= right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        *swaps++
        k++
    }

    for i < len1 {
        arr[k] = left[i]
        i++
        k++
        *swaps++
    }

    for j < len2 {
        arr[k] = right[j]
        j++
        k++
        *swaps++
    }
}

func timSort(arr []int, comparisons, swaps *int) {
    n := len(arr)
    for i := 0; i < n; i += RUN {
        insertionSort(arr, i, min(i+RUN-1, n-1), comparisons, swaps)
    }

    for size := RUN; size < n; size = 2 * size {
        for left := 0; left < n; left += 2 * size {
            mid := min(left+size-1, n-1)
            right := min(left+2*size-1, n-1)
            if mid < right {
                merge(arr, left, mid, right, comparisons, swaps)
            }
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func (t TimSort) Name() string {
    return "Tim Sort"
}