package arrutil

// matchStringArray will compare two string arrays regardless of their order.
func matchStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	diff := make(map[string]bool)

	for _, dim := range a {
		diff[dim] = true
	}

	for _, dim := range b {
		if !diff[dim] {
			return false
		}
	}

	return true
}
