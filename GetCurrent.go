package directory

import (
	"os"
)

// GetCurrent - simple wrapper for os.Getwd()
func GetCurrent() string {
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return p
}
