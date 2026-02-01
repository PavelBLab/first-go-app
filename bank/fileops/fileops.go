package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetFLoatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, nil
	}

	valueText := strings.TrimSpace(string(data))
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}