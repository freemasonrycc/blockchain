package types

// Check if the key is valid(key is not nil)
func AssertValidKey(key []byte) {
	if len(key) == 0 {
		panic("key is nil")
	}
}

// Check if the value is valid(value is not nil)
func AssertValidValue(value []byte) {
	if value == nil {
		panic("value is nil")
	}
}
