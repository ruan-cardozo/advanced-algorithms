package algorithms

type MergeSort struct{}

func (m MergeSort) Sort(arr []int) ([]int, int, int) {
	comparisons, swaps := 0, 0
	sortedArray := mergeSort(arr, &comparisons, &swaps)
	return sortedArray, comparisons, swaps
}

func mergeSort(arr []int, comparisons, swaps *int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid], comparisons, swaps)
	right := mergeSort(arr[mid:], comparisons, swaps)

	return Merge(left, right, comparisons, swaps)
}

func Merge(left, right []int, comparisons, swaps *int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		*comparisons++
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
			*swaps++
		} else {
			result[k] = right[j]
			j++
			*swaps++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		*swaps++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		*swaps++
		k++
	}

	return result
}

func (m MergeSort) Name() string {
	return "Merge Sort"
}