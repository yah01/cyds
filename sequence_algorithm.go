package cyds

func BinarySearch(slice []interface{}, condition func(i int) bool) int {
	var (
		l   = 0
		r   = len(slice) - 1
		mid = 0
	)

	for l < r {
		mid = (l + r) >> 1
		if condition(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}
