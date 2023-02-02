package opfile

import (
	"os"
)

// CreateDirIfNotExist create dir recursion
func CreateDirIfNotExist(dir string) error {
	if CheckPathIfNotExist(dir) {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}

// CheckPathIfNotExist if file exist return true
func CheckPathIfNotExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// RemoveFile whatever file exist or not
func RemoveFile(path string) {
	_ = os.Remove(path)
}
