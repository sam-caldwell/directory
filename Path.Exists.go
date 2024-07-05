package directory

// Exists - Return whether the path exists.
//
// go:inline
func (p *Path) Exists() bool {
	return Exists(string(*p))
}
