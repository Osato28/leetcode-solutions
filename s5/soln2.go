// Learning algo 1, deliberately inefficient: implementing brute-force algo from editorial
func longestPalindrome(s string) string {
	for length := len(s); length > 0; length-- {
		for start := 0; start <= len(s)-length; start++ {
			if check(start, start+length, s) {
				return s[start : start+length]
			}
		}
	}
	return ""
}

func check(i int, j int, s string) bool {
	left := i
	right := j - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}