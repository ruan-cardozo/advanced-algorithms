package algorithms

type BubbleSortImproved struct{}

func (b BubbleSortImproved) Sort(arr []int) ([]int, int, int) {
	n := len(arr)
	comparisons, swaps := 0, 0
	sorted := make([]int, n)
	copy(sorted, arr)

	swapped := true
	for i := 0; i < n-1 && swapped; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			comparisons++
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				swaps++
				swapped = true
			}
		}
	}
	return arr, comparisons, swaps
}

func (b BubbleSortImproved) Name() string {
	return "Bubble Sort Improved"
}