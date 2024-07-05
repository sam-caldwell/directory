package directory

// Exists - Return boolean value if directory exists
//
//go:inline
func Exists(name string) bool {
	return Existsp(&name)
}
