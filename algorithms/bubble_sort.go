package algorithms

type BubbleSort struct{}

func (b BubbleSort) Sort(arr []int) ([]int, int, int) {
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

func (b BubbleSort) SortWithProgress(arr []int) ([]int, int, int) {
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



func (b BubbleSort) Name() string {
    return "Bubble Sort"
}