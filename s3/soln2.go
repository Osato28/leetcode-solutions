func lengthOfLongestSubstring(s string) int {
	// Hashmap v2, but with complexity O(n^2) instead of O(n)
	var maxLength int = 0
	hashmap := make(map[rune]int)
	for i, char := range s {
		hashmap[char] = i
		for charInMap, position := range hashmap {
			if i-position > len(hashmap) {
				delete(hashmap, charInMap)
			}
		}
		if maxLength < len(hashmap) {
			maxLength = len(hashmap)
		}
	}
	return len(hashmap)
}