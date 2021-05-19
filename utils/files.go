package utils

import (
	"os"
)

func OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)

	return file, err
}