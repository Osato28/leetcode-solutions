func findPeakElement(nums []int) int {
	numsLen := len(nums)

	// process edge cases before doing binary search, both for speed and to keep the algorithm lean with as few conditionals as possible
	if numsLen == 0 {
		return -1
	} else if numsLen == 1 {
		return 0
	} else {
		// check edges for peaks
		if nums[0] > nums[1] {
			return 0
		}
		if nums[numsLen-1] > nums[numsLen-2] {
			return numsLen - 1
		}
	}

	// binary search for the local maximum
	left, right := 0, numsLen-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		}

		if nums[mid] < nums[mid+1] {
			left = mid
		} else if nums[mid] < nums[mid-1] {
			right = mid
		}
	}
	return -1
}