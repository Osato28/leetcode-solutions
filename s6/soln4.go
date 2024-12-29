// Second intuition w/optimization:
// use an empirically derived formula to append strings to result string in place, without separating it into rows
// use strings.Builder for concatenation instead of converting byte to string and then concatenating string

// Lesson learned: in Go, use strings.Builder when you need to do a lot of operations on strings

// general case
// period 1: (numRows - 1) * 2
// period N: (numRows - 1 - N) * 2; 2 * N

func convert(s string, numRows int) string {
	var sb strings.Builder

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
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}