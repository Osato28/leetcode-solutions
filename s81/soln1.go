func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	pivot := findRotationPivot(nums)
	var searchResult bool

	if pivot == 0 {
		return binarySearch(nums, target, 0, len(nums)-1)
	} else {
		searchResult = binarySearch(nums, target, 0, pivot-1)
		if !searchResult {
			searchResult = binarySearch(nums, target, pivot, len(nums)-1)
		}
		return searchResult
	}
}

func binarySearch(nums []int, target int, left int, right int) bool {
	for left <= right {
		median := (left + right) / 2

		if nums[median] == target {
			return true
		} else if nums[median] < target {
			left = median + 1
		} else if nums[median] > target {
			right = median - 1
		}
	}
	return false
}

func findRotationPivot(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	if nums[len(nums)-1] > nums[0] {
		return 0
	}
	pivot := 1
	for {
		if pivot >= len(nums) {
			return 0
		}
		if nums[pivot] < nums[pivot-1] {
			break
		}
		pivot++
	}
	return pivot
}