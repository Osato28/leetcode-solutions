// Naive implementation of Expand-Around-Center, fails on test case 141/142
func longestPalindrome(s string) string {
	// max palindrome length
	lpLen := 1
	lpStr := string(s[:1])
	for i, _ := range s {
		if !enoughSpaceForPalindrome(lpLen, i, len(s)) {
			continue
		}
		currLp := findLongestPalindrome(s, i, lpLen)
		if lpLen < len(currLp) {
			lpStr = currLp
			lpLen = len(lpStr)
		}
	}
	return lpStr
}

func findLongestPalindrome(s string, i int, lpLen int) string {
	longestPalindrome := string(s[i])
	isPalindrome := false
	// method: expand from center
	// the center is at index i, like this: A, Aa, bAb, bAab, cbAbc
	for pLen := lpLen; pLen <= len(s); pLen++ {
		enoughSpace := enoughSpaceForPalindrome(pLen, i, len(s))
		isEven := pLen%2 == 0
		// Check whether there is enough space for the palindrome on both sides
		if enoughSpace {
			for dist := 1; dist <= pLen/2; dist++ {
				if isEven {
					// even length
					isPalindrome = (s[i-dist+1] == s[i+dist])
					if isPalindrome {
						longestPalindrome = s[i-dist+1 : i+dist+1]
					}
				} else {
					// odd length
					isPalindrome = (s[i-dist] == s[i+dist])
					if isPalindrome {
						longestPalindrome = s[i-dist : i+dist+1]
					}
				}
				if !isPalindrome {
					break
				}
			}
		}
		if !enoughSpace {
			break
		}
	}
	return longestPalindrome
}

func enoughSpaceForPalindrome(pLen int, i int, sLen int) bool {
	// empirically derived formulae, should be fine but might be a source of bugs
	esLeft := i+1 >= int(math.Ceil(float64(pLen)/2))
	esRight := (sLen-i > pLen/2)

	return esLeft && esRight
}