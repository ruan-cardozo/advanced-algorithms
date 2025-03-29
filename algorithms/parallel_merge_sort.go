package algorithms

const THRESHOLD = 500000

type ParallelMergeSort struct {
	threshold int
}

type sortResult struct {
	array       []int
	comparisons int
	swaps       int
}

func (ms ParallelMergeSort) Name() string {
	return "Parallel Merge Sort"
}

func (ms ParallelMergeSort) Sort(arr []int) ([]int, int, int) {
	return ms.SortParallel(arr)
}

func (ms ParallelMergeSort) SortParallel(arr []int) ([]int, int, int) {
	if ms.threshold == 0 {
		ms.threshold = THRESHOLD
	}

	if len(arr) <= ms.threshold {
		return MergeSort{}.Sort(arr)
	}

	mid := len(arr) / 2

	leftChan := make(chan sortResult)
	rightChan := make(chan sortResult)

	go func() {
		left, c, s := ms.SortParallel(arr[:mid])
		leftChan <- sortResult{array: left, comparisons: c, swaps: s}
	}()

	go func() {
		right, c, s := ms.SortParallel(arr[mid:])
		rightChan <- sortResult{array: right, comparisons: c, swaps: s}
	}()

	leftResult := <-leftChan
	rightResult := <-rightChan

	result := make([]int, len(arr))
	comparisons, swaps := simpleMerge(leftResult.array, rightResult.array, result)

	return result,
		comparisons + leftResult.comparisons + rightResult.comparisons,
		swaps + leftResult.swaps + rightResult.swaps
}

func simpleMerge(left, right []int, result []int) (int, int) {
	comparisons, swaps := 0, 0
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		comparisons++
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		swaps++
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
		swaps++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
		swaps++
	}

	return comparisons, swaps
}
