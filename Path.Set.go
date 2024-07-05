package directory

// Set - Setter: store a new string (v) to the Path.
//
// go:inline
func (p *Path) Set(v string) {
	p.Setp(&v)
}
