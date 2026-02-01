package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManger) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		fmt.Println("Error reading file:", err)
		return nil, errors.New("failed to read file")
	}

	file.Close()

	return lines, nil
}

func (fm *FileManger) WriteResults(data interface{}) error { // interface{} or any means that any data can be passed
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return err
	}
	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

func New(inputFilePath, outputFilePath string) *FileManger {
	return &FileManger{
		inputFilePath,
		outputFilePath,
	}
}
