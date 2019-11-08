package main

import "os"

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
// https://stackoverflow.com/a/12518877
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
