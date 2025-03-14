package algorithms

type SelectionSort struct{}

func (s SelectionSort) Sort(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	for i := 0; i < len(sorted)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(sorted); j++ {
			if sorted[j] < sorted[minIdx] {
				minIdx = j
			}
		}
		sorted[i], sorted[minIdx] = sorted[minIdx], sorted[i]
	}
	return sorted
}
