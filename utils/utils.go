package utils

import (
	"fmt"
	"strconv"
)

func SafeAtoi(value string) int {
	if value == "" {
		return 0
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Error converting to int: %v\n", err)
		return 0
	}
	return result
}

func SafeParseFloat(value string, precision int) float64 {
	if value == "" {
		return 0
	}
	result, err := strconv.ParseFloat(value, precision)
	if err != nil {
		fmt.Printf("Error converting to float: %v", err)
		return 0
	}

	return result
}
