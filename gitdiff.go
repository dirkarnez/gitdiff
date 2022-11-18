package gitdiff

import (
	"os"
)

func GetLeft() (string, error) {
	dat, err := os.ReadFile(os.Args[1])
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func GetRight() (string, error) {
	dat, err := os.ReadFile(os.Args[2])
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
