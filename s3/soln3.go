func lengthOfLongestSubstring(s string) int {
	// Utilizing sliding window on a set (hashmap)
	maxLength := 0
	left := 0
	runeSet := make(map[rune]int)
	var runeArr []rune = []rune(s)
	for right, newChar := range runeArr {
		for {
			_, isPresent := runeSet[newChar]
			if isPresent {
				delete(runeSet, runeArr[left])
				left++
			} else {
				break
			}
		}
		runeSet[newChar] = right

		length := right - left + 1
		if maxLength < length {
			maxLength = length
		}
	}
	return maxLength
}