package util

func Contain[T comparable](elems []T, v T) bool {
	for _, e := range elems {
		if e == v {
			return true
		}
	}
	return false
}
