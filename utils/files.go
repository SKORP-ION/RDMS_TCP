package utils

import (
	"bufio"
	"bytes"
	"os"
)

func ReadFile(path string) (bytes.Buffer, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return bytes.Buffer{}, err
	}

	buffer := bytes.Buffer{}
	scan := bufio.NewScanner(file)

	for scan.Scan() {
		buffer.Write(scan.Bytes())
	}

	return buffer, nil
}