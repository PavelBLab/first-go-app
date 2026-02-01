package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for strIndex, str := range strings {
		floatValue, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to convert string to float at line %d: %v", strIndex+1, err)
		}

		floats[strIndex] = floatValue
	}

	return floats, nil
}
