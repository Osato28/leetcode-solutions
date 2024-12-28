func searchMatrix(matrix [][]int, target int) bool {
	top, down := 0, len(matrix)-1
	for top <= down {
		median := (top + down) / 2
		row := matrix[median]

		if mightContainNumber(row, target) {
			return binarySearchInRow(row, target)
		}

		if target < row[0] {
			if median == top {
				break
			}
			down = median - 1
		} else if target > row[len(row)-1] {
			if median == down {
				break
			}
			top = median + 1
		}
	}
	return false
}

func mightContainNumber(row []int, target int) bool {
	if row[0] <= target && row[len(row)-1] >= target {
		return true
	}
	return false
}

func binarySearchInRow(row []int, target int) bool {
	left, right := 0, len(row)-1
	for left <= right {
		median := (left + right) / 2

		if row[median] == target {
			return true
		}

		if row[median] < target {
			left = median + 1
		} else if row[median] > target {
			right = median - 1
		}
	}
	return false
}