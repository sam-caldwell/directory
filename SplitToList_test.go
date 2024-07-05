package directory

import (
	"testing"
)

func TestSplitToList(t *testing.T) {
	t.Run("Happy: Split valid path (/etc/passwd)", func(t *testing.T) {
		result := SplitToList("/etc/passwd")
		if len(result) != 2 {
			for i, e := range result {
				t.Logf("result %d: '%s'", i, e)
			}
			t.Fatalf("expect %d elements, got %d elements: %v", 2, len(result), result)
		}
	})

	t.Run("Happy: Split valid path (/etc/passwd)", func(t *testing.T) {
		const testPath = "/tmp/1/2/3/4/5"
		result := SplitToList(testPath)
		if len(result) != 6 {
			for i, e := range result {
				t.Logf("result %d: '%s'", i, e)
			}
			t.Fatalf("expect %d elements, got %d elements: %v", 2, len(result), result)
		}
	})

}
