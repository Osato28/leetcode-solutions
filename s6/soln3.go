// Second intuition:
// use an empirically derived formula to append strings to result string in place, without separating it into rows

// General case
// period 1: (numRows - 1) * 2; 0
// period 2: (numRows - 2) * 2; 2
// period n: (numRows - 1 - n) * 2; 2 * n
// period N: 0; (numRows - 1) * 2

// Works faster than the first intuition, but still slower than 80% of solutions

func convert(s string, numRows int) string {
	rows := make([]string, numRows)
	result := ""
	if numRows == 1 {
		return s
	}
	for n := 0; n < numRows; n++ {
		even := false
		period := 0
		for i := n; i < len(s); i += period {
			if even {
				period = n * 2
			} else {
				period = (numRows - 1 - n) * 2
			}
			even = !even
			if period == 0 {
				continue
			}
			rows[n] += string(s[i])
			result += string(s[i])
		}
	}
	return result
}