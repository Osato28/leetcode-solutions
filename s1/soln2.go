// Using a hash table AKA Go map for O(n) complexity
func twoSum(nums []int, target int) []int {
	// Fill in a map in n steps
	m := make(map[int]int)
	for i, val := range nums {
		m[val] = i
	}

	// Find complement for every number (n steps)
	for i, val := range nums {
		complementVal := target - val
		j, ok := m[complementVal]
		if ok && i != j {
			// If complement found, return
			return []int{i, j}
		}
	}

	return []int{}
}