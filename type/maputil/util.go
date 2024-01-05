package maputil

// CopyFrom creates a new map from the original map, essentially duplicating it with the same key-value pairs.
func CopyFrom(arr map[int]int) map[int]int {
	arrayNew := make(map[int]int)
	for key, value := range arr {
		arrayNew[key] = value
	}
	return arrayNew
}
