// That was a pretty interesting thing to do with binary search.
func mySqrt(x int) int {
	left := 0
	right := x

	for {
		median := int(math.Floor(float64(left+right) / 2))

		if left-right == 1 {
			return right
		} else if right < left {
			break
		}

		square := median * median
		if square < x {
			left = median + 1
		} else if square > x {
			right = median - 1
		} else {
			return median
		}
	}

	// If the program returns -1, the algorithm has failed to find a square root, so it's borked
	return -1
}