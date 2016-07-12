package jsonfile

import "os"

// FileExists - function to test the existance of a file
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
