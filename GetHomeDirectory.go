package directory

import "os/user"

// GetHomeDirectory - returns the current user's home directory
func GetHomeDirectory() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
