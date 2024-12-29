// Second solution, mildly optimized by inserting elements from a smaller array into a larger one & using native Go slices.Insert to insert values
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Insert elements from a smaller array into a larger one
	var mergedNums []int
	if len(nums1) > len(nums2) {
		mergedNums = mergeArrays(nums1, nums2)
	} else {
		mergedNums = mergeArrays(nums2, nums1)
	}

	return findMedianInArray(mergedNums)
}

func mergeArrays(targetArray []int, sourceArray []int) []int {
	if len(targetArray) == 0 {
		return sourceArray
	}

	resultArr := targetArray
	for _, value := range sourceArray {
		// mergesort: find insertion point via binary search O(logm) and then insert it, repeat for n elements
		insIndex := findInsertionPoint(resultArr, value)
		resultArr = slices.Insert(resultArr, insIndex, value)
	}
	return resultArr
}

func findInsertionPoint(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid
		}
	}
	if nums[right] >= target {
		return right
	} else {
		return right + 1
	}
}

func findMedianInArray(arr []int) float64 {
	mid := (len(arr) - 1) / 2
	// if len is even, average two middle elements
	// otherwise return middle element
	if len(arr)%2 == 0 {
		return float64(arr[mid]+arr[mid+1]) / 2
	} else {
		return float64(arr[mid])
	}
}