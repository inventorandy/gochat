package ptr

// Bool creates a boolean pointer
func Bool(value bool) *bool {
	rtn := new(bool)
	*rtn = value
	return rtn
}

// String creates a string pointer
func String(value string) *string {
	rtn := new(string)
	*rtn = value
	return rtn
}
