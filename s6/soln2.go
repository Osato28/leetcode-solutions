// Attempt to speed up v1 by excluding modulo operations
// But the algorithm is still much slower than most

func convert(s string, numRows int) string {
	rows := make([]string, numRows)

	if numRows == 1 {
		return s
	}

	rowIndex := 0
	dir := -1
	for _, char := range s {
		rows[rowIndex] += string(char)

		if rowIndex == 0 || rowIndex == numRows-1 {
			dir = -dir
		}
		rowIndex += dir
	}

	var result string = ""
	for _, row := range rows {
		fmt.Println(row)
		result += row
	}
	return result
}