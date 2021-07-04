package ptr

// Bool creates a boolean pointer
func Bool(value bool) *bool {
	rtn := new(bool)
	*rtn = value
	return rtn
}
