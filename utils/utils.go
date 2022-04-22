package utils

func Clamp(x, low, high int) (int, bool) {
	if x > low && x < high {
		return x, false
	} else if x <= low {
		return low, true
	} else {
		return high, true
	}
}
