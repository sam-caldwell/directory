package directory

// Setp - Setter: store a new string (*v) to the Path.
//
// go:inline
func (p *Path) Setp(v *string) {
	//ToDo: validate path is valid/sanitized string
	*p = Path(*v)
}
