package fileutils

import (
	"io/ioutil"
)

// ReadFile reads file
func ReadFile(filepath string) (string, error) {
	var dat, err = ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
