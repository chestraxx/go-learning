package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("error reading file")
	}

	return lines, nil
}

func WriteFileJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("error creating file")
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return errors.New("error encoding data")
	}

	return nil
}
