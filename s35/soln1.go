// First implementation of binary search, doesn't work properly
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		median := int(math.Floor(float64(left+right) / 2.0))

		if left == len(nums) || right-left == 1 {
			fmt.Printf("found place for target, left %d, median %d, right %d\n", left, median, right)
			return left
		}

		if left > right {
			fmt.Printf("failed, left %d, median %d, right %d\n", left, median, right)
			break
		}

		fmt.Printf("iterating, left %d, median %d, right %d\n", left, median, right)
		if nums[median] > target {
			right = median - 1
		} else if nums[median] < target {
			left = median + 1
		} else {
			fmt.Printf("found target, left %d, median %d, right %d\n", left, median, right)
			return median
		}
	}
	// If the program returns -1, it failed to find anything
	return -1
}