package algorithms

type HeapSort struct{}

func (h HeapSort) heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		h.heapify(arr, n, largest)
	}
}

func (h HeapSort) Sort(arr []int) ([]int, int, int) {
	n := len(arr)
	comparisons, swaps := 0, 0

	for i := n/2 - 1; i >= 0; i-- {
		h.heapify(arr, n, i)
		comparisons++
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		swaps++

		h.heapify(arr, i, 0)
		comparisons++
	}
	return arr, comparisons, swaps
}

func (h HeapSort) Name() string {
	return "Heap Sort"
}
