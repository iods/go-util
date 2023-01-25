package maputil

// MapStringInt declares a wrapper for the `map[string]int`
type MapStringInt map[string]int

// NewMapStringIntSlice creates a new ...
func NewMapStringIntSlice(slice []string) MapStringInt {
	mapstringint := MapStringInt{}
	for i, s := range slice {
		mapstringint[s] = i
	}
	return mapstringint
}

// Set sets the value of `v` to `k`
func (m MapStringInt) Set(k string, v int) {
	m[k] = v
}
