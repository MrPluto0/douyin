package arr

import "sort"

func IsContain(target string, array []string) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == target {
			return true
		}
	}
	return false
}

// Don't use. This func has lower efficiency.
func _IsContain(target string, array []string) bool {
	index := sort.SearchStrings(array, target)
	if index < len(array) && array[index] == target {
		return true
	}
	return false
}
