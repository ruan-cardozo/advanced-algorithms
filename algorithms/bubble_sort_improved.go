package algorithms

type BubbleSortImproved struct{}

func (b BubbleSortImproved) Sort(arr []int) []int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)

	swapped := true
	for i := 0; i < n-1 && swapped; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				swapped = true
			}
		}
	}
	return sorted
}
