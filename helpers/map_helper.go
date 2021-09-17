package helpers

func SliceIntContainsValue(m []int, v interface{}) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}
	return false
}
