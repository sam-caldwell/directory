package directory

// String - Getter: return string path
//
// go:inline
func (p *Path) String() string {
	return (string)(*p)
}
