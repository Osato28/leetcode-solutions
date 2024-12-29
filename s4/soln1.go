// First naive solution, complexity n*logn
// Technically works, but is in the bottom 5% in terms of performance
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedNums := mergeArrays(nums1, nums2)
	return findMedianInArray(mergedNums)
}

func mergeArrays(targetArray []int, sourceArray []int) []int {
	if len(targetArray) == 0 {
		return sourceArray
	}

	resultArr := targetArray
	for _, value := range sourceArray {
		// insert each value from sourceArray into the right point to keep the merged array sorted
		insIndex := findInsertionPoint(resultArr, value)
		resultArr = append(resultArr[:insIndex], append([]int{value}, resultArr[insIndex:]...)...)
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