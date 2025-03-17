package algorithms

type QuickSort struct{}

func (q QuickSort) Sort(arr []int) ([]int, int, int) {
	var comparisons, swaps int

	partition := func(arr []int, low, high int) ([]int, int) {
		pivot := arr[high]
		i := low
		for j := low; j < high; j++ {
			comparisons++
			if arr[j] < pivot {
				arr[i], arr[j] = arr[j], arr[i]
				swaps++
				i++
			}
		}
		arr[i], arr[high] = arr[high], arr[i]
		swaps++
		return arr, i
	}

	var quickSort func([]int, int, int) []int
	quickSort = func(arr []int, low, high int) []int {
		if low < high {
			var p int
			arr, p = partition(arr, low, high)
			arr = quickSort(arr, low, p-1)
			arr = quickSort(arr, p+1, high)
		}
		return arr
	}

	sortedArr := quickSort(arr, 0, len(arr)-1)
	return sortedArr, comparisons, swaps
}