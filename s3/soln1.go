// Doesn't work out: the map turns out to be a subsequence, not a substring
func lengthOfLongestSubstring(s string) int {
	// First try: use hashmap
	var hashmap map[rune]int = make(map[rune]int)
	for i, char := range s {
		fmt.Println(char)
		j, alreadyInMap := hashmap[char]
		if alreadyInMap {
			if i-j > len(hashmap)-1 {
				clear(hashmap)
			} else {
				delete(hashmap, char)
			}
		}
		hashmap[char] = i
	}
	fmt.Println(hashmap)
	return len(hashmap)
}