package algorithms

var cache = make(map[int]int)

func BinarySearch(sortedArray []int, low, high, target int) int {

	if val, ok := cache[target]; ok {
		return val
	}

	if low > high {
		cache[target] = -1
		return -1
	}

	mid := low + (high-low)/2

	if sortedArray[mid] == target {

		if mid == 0 || sortedArray[mid-1] != target {
			cache[target] = mid
			return mid
		}
		return BinarySearch(sortedArray, low, mid-1, target)
	} else if sortedArray[mid] > target {
		return BinarySearch(sortedArray, low, mid-1, target)
	} else {
		return BinarySearch(sortedArray, mid+1, high, target)
	}
}