package directory

// Bytes - Getter: return a []byte path
//
// go:inline
func (p *Path) Bytes() []byte {
	return ([]byte)(*p)
}
