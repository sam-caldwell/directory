package directory

import (
	"path/filepath"
	"strings"
)

// SplitToList - Given a path (possibly ending in a file name) return a list of only the directories.
func SplitToList(path string) []string {
	if !strings.HasSuffix(path, string(filepath.Separator)) {
		path += string(filepath.Separator)
	}
	thisDir, _ := filepath.Split(path)
	list := strings.Split(thisDir, string(filepath.Separator))
	list = RemoveEmptyElements(&list)
	return list
}

// RemoveEmptyElements removes empty elements from a []string.
// ToDo: move this to lists package
func RemoveEmptyElements(input *[]string) []string {
	var result []string
	for _, str := range *input {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
