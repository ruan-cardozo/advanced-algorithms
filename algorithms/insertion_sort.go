package algorithms

type InsertionSort struct{}

func (i InsertionSort) Sort(arr []int) ([]int, int, int) {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	comparisons, swaps := 0, 0

	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 {
			comparisons++
			if sorted[j] > key {
				sorted[j+1] = sorted[j]
				swaps++
				j--
			} else {
				break
			}
		}
		sorted[j+1] = key
	}
	return sorted, comparisons, swaps
}
