// O(n^2) time complexity, does not satisfy the followup question
func twoSum(nums []int, target int) []int {
	for i, val1 := range nums {
		if i == len(nums)-1 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			val2 := nums[j]
			if val1+val2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}