package algorithms

func BinarySearch(sortedArray []int, low, high, target int) int {

	cache := make(map[int]int)

	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if sortedArray[mid] == target {
		// Check if it's the first occurrence
		if mid == 0 || sortedArray[mid-1] != target {
			cache[mid] = target
			return mid
		}
		return BinarySearch(sortedArray, low, mid-1, target)
	} else if sortedArray[mid] > target {
		return BinarySearch(sortedArray, low, mid-1, target)
	} else {
		return BinarySearch(sortedArray, mid+1, high, target)
	}
}