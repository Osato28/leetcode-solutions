// V2, fixed after modeling the algo with physical props
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		median := int(math.Floor(float64(left+right) / 2.0))

		if left-right == 1 {
			return left
		} else if left > right {
			break
		}

		if nums[median] > target {
			right = median - 1
		} else if nums[median] < target {
			left = median + 1
		} else {
			return median
		}
	}
	// If the program returns -1, it failed to find a place, meaning the algorithm is borked somehow
	return -1
}