// Do binary search in a sorted array
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for {
		median := int(math.Floor(float64(left+right) / 2))

		if nums[median] < target {
			left = median + 1
		} else if nums[median] > target {
			right = median - 1
		} else {
			return median
		}

		if left > right {
			break
		}
	}

	return -1
}

// First (naive) solution: simply rotate the array back to a sorted state before searching and then rotate the search result
func search(nums []int, target int) int {
	arrLength := len(nums)

	// find rotation pivot if it exists (located to the right of the point of rotation or at 0 if there is none)
	rotPivot := 0
	for i := 1; i < arrLength; i++ {
		if nums[i] < nums[i-1] {
			rotPivot = i
		}
	}

	// unrotate array
	var originalNums []int
	originalNums = make([]int, arrLength)
	for i := rotPivot; i < arrLength+rotPivot; i++ {
		j := i % arrLength
		originalNums[i-rotPivot] = nums[j]
	}

	// do binary search
	resultOfSearch := binarySearch(originalNums, target)

	// rotate the resulting value
	if resultOfSearch > -1 {
		return (resultOfSearch + rotPivot) % arrLength
	}
	return -1
}