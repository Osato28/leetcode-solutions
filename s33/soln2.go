
func findIndexOfMinimum(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		median := (left + right) / 2
		if nums[median] > nums[right] {
			left = median + 1
		} else {
			right = median
		}
	}

	return left
}
func binarySearch(nums []int, target int, start int, end int) int {
	left := start
	right := end

	for left <= right {
		median := (left + right) / 2
		if nums[median] < target {
			left = median + 1
		} else if nums[median] > target {
			right = median - 1
		} else {
			return median
		}
	}

	return -1
}

// Second solution: search in the rotated array by searching the two parts separately
func search(nums []int, target int) int {
	minimumIndex := findIndexOfMinimum(nums)

	// If the target is greater/equal than the value at index 0, search in 0..minimumIndex-1
	// Otherwise search in minimumIndex..len(nums)-1
	if minimumIndex > 0 && nums[0] <= target {
		return binarySearch(nums, target, 0, minimumIndex-1)
	} else {
		return binarySearch(nums, target, minimumIndex, len(nums)-1)
	}
}