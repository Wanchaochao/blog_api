package mathutil

// PowInt is int type of math.Pow function.
func PowInt(x int, y int) int {
	if y <= 0 {
		return 1
	} else {
		if y%2 == 0 {
			sqrt := PowInt(x, y/2)
			return sqrt * sqrt
		} else {
			return PowInt(x, y-1) * x
		}
	}
}

func AbsInt(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

func MinInt(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt64(a int64, b int64) int64 {
	if a > b {
		return b
	} else {
		return a
	}
}

func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}