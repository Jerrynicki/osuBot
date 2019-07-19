package util

// CheckForErrors panics when the error is not nil
func CheckForErrors(e error) {
	if e != nil {
		panic(e)
	}
}
