package algorithms

type SelectionSort struct{}

func (s SelectionSort) Sort(arr []int) ([]int, int, int) {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	comparisons, swaps := 0, 0

	for i := 0; i < len(sorted)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(sorted); j++ {
			comparisons++
			if sorted[j] < sorted[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			sorted[i], sorted[minIdx] = sorted[minIdx], sorted[i]
			swaps++
		}
	}
	return sorted, comparisons, swaps
}

func (s SelectionSort) Name() string {
	return "Selection Sort"
}