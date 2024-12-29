// Learning algo 2: implementing DP algo from editorial, by memory
func longestPalindrome(s string) string {
	n := len(s)
	ans := []int{0, 0}

	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	// Length 1 palindromes
	for i := 0; i < n; i++ {
		// Single characters are palindromes by definition
		isPalindrome[i][i] = true
	}
	// Length 2 palindromes
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			isPalindrome[i][i+1] = true
			ans = []int{i, i + 1}
		}
	}
	// Length 3+ palindromes
	for diff := 2; diff < n; diff++ {
		for i := 0; i < n-diff; i++ {
			j := i + diff
			if s[i] == s[j] && isPalindrome[i+1][j-1] {
				isPalindrome[i][j] = true
				ans = []int{i, j}
			}
		}
	}

	return s[ans[0] : ans[1]+1]
}
