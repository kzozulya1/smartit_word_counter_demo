package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
)

// ScanDir scans dir for files
func ScanDir(dir string) ([]string, error) {
	var files []string

	//check dir exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return files, fmt.Errorf("dir %s doesn't exist", dir)
	}

	var err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		//skip dirs
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	return files, err
}
