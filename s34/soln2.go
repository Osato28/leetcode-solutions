// True O(logn) implementation: find upper and lower bounds in two binary searches
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	upper := findUpperBound(nums, target)
	lower := findLowerBound(nums, target)
	return []int{lower, upper}
}

func findUpperBound(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// This is very inelegant. Is there any way to rewrite the logic into a more straightforward form?
	if right == 0 {
		return -1
	}
	if right == len(nums)-1 && nums[right] == target {
		return right
	}
	if nums[right-1] == target {
		// If any elements with the target value are present,
		return right - 1
	}
	return -1
}

func findLowerBound(nums []int, target int) int {
	// This algorithm doesn't work in arrays of length 1
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		// If any elements with the target value are present, return the index of the left
		return left
	}
	return -1
}