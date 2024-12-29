// Naive intuition:
// the letters at the top and bottom are occurring every numRows-1 chars.

// So top letters all have an index such that i % 2*(numRows-1) == 0
// Bottom letters have an index such that i % (numRows-1) == 0 && i % 2*(numRows-1) != 0
// The ones on nth row have an index such that i % (numRows-1) == n

// Initial intuition: make numRows strings (one for each row), push letters onto them
// After last letter was processed, concatenate them

func convert(s string, numRows int) string {
	rows := make([]string, numRows)

	if numRows == 1 {
		return s
	}
	for i, char := range s {
		remainder := i % (numRows - 1)
		dividend := i / (numRows - 1)

		if dividend%2 == 0 {
			rows[remainder] += string(char)
		} else {
			rows[numRows-1-remainder] += string(char)
		}
	}

	var result string = ""
	for _, row := range rows {
		fmt.Println(row)
		result += row
	}
	return result
}