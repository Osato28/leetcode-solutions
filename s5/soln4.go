// Learning algo 3: implementing expand-from-center algo, half-optimized
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""

	expand := func(i int, j int) {
		left, right := i, j

		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			}
			if right-left+1 > len(ans) {
				ans = s[left : right+1]
			}
			left--
			right++
		}
	}

	for i := 0; i < n; i++ {
		if i < n-(len(ans)/2)-1 && s[i] == s[i+1] {
			expand(i, i+1)
		}
		if i < n-(len(ans)/2) {
			expand(i, i)
		}
	}

	return ans
}
