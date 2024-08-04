package directory

import (
	"os/user"
	"testing"
)

func TestGetHomeDirectory(t *testing.T) {
	if expectedUser, err := user.Current(); err != nil {
		panic(err)
	} else {
		if homeDir, err := GetHomeDirectory(); err != nil {
			t.Errorf("Unexpected error: %v", err)
		} else {
			if homeDir != expectedUser.HomeDir {
				t.Errorf("Expected home directory to be /home/user, got %s", homeDir)
			}
		}
	}
}
