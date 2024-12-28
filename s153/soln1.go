// Extremely inelegant. Still can't think in terms of binary search, I end up having to do trial and error with every new problem
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	// this implementation of the algorithm gets stuck when n is even (at left,median being at the maximum and right being at the minimum), so it needs an additional check for whether it's stuck
	moved := true
	for moved && (left < right) {
		moved = false
		median := (left + right) / 2

		// left moves to the maximum (to the left of the minimum)
		if (nums[left] < nums[median]) && (nums[right] < nums[median]) {
			left = median
			moved = true
		}

		// Once median has traveled to the unrotated portion of the array, right travels towards it and thus towards the minimum
		if nums[right] > nums[median] {
			right = median
			moved = true
		}
	}
	if right == 0 {
		return nums[0]
	}
	return nums[right]
}