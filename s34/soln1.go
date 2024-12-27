// Naive implementation: find any index of the target value in logn steps, then find edges of the target values in k steps
func searchRange(nums []int, target int) []int {
	// find any element within the range where all elements have target value
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	midTargetIndex := binarySearch(nums, target)
	fmt.Println(midTargetIndex)
	if midTargetIndex == -1 {
		return []int{-1, -1}
	}

	// traverse in k steps to find edges of the target value
	var min, max, i int
	for i = midTargetIndex; i >= 0; i-- {
		if nums[i] == target {
			min = i
		} else {
			break
		}
	}
	for i = midTargetIndex; i < len(nums); i++ {
		if nums[i] == target {
			max = i
		} else {
			break
		}
	}
	return []int{min, max}
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

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