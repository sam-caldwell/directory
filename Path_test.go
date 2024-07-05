package directory

import (
	"testing"
)

func TestPath(t *testing.T) {
	const test = "foo"
	s := string(Path(string(test)))
	if s != test {
		t.Fatal("Path should cast to string without error")
	}
}
