package uti

func UnorderMatch[T comparable](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	m1, m2 := map[T]bool{}, map[T]bool{}
	for _, c := range arr1 {
		m1[c] = true
	}
	for _, c := range arr2 {
		m2[c] = true
	}
	for _, c := range arr1 {
		_, ok := m2[c]
		if !ok {
			return false
		}
	}
	for _, c := range arr2 {
		_, ok := m1[c]
		if !ok {
			return false
		}
	}
	return true
}
