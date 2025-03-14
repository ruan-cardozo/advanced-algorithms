package algorithms

type InsertionSort struct{}

func (i InsertionSort) Sort(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j] > key {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	return sorted
}
