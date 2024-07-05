package directory

import "fmt"

// NewPath - Return a new path with a given string or []byte value.
//
// Return a new Path object (string) and nil error if the path exists.
// Return nil if the path must exist and does not.
func NewPath[T string | []byte](path T, mustExist bool) (*Path, error) {
	var o Path
	if mustExist {
		if !Exists(string(path)) {
			return nil, fmt.Errorf("path must exist but not found")
		}
	}
	o.Set(string(path))
	return &o, nil
}
