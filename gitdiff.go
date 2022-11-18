package gitdiff

import (
	"io/fs"
	"io/ioutil"
	"os"
)

func GetLeftFile() (string, error) {
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func GetRightFile() (string, error) {
	dat, err := os.ReadFile(os.Args[2])
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func GetLeftDirRecursive() ([]fs.FileInfo, error) {
	return ioutil.ReadDir(os.Args[1])
}

func GetRightDirRecursive() ([]fs.FileInfo, error) {
	return ioutil.ReadDir(os.Args[2])
}
