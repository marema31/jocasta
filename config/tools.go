package config

import (
	"math"
	"strconv"
	"strings"
)

func sizeFromString(str string) (int, error) {
	var power float64
	var strValue string

	switch {
	case strings.Contains(str, "TB"):
		power = 4
		strValue = str[:len(str)-2]
	case strings.Contains(str, "GB"):
		power = 3
		strValue = str[:len(str)-2]
	case strings.Contains(str, "MB"):
		power = 2
		strValue = str[:len(str)-2]
	case strings.Contains(str, "KB"):
		power = 1
		strValue = str[:len(str)-2]
	case strings.Contains(str, "B"):
		power = 0
		strValue = str[:len(str)-1]
	default:
		power = 0
		strValue = str
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}

	return value * int(math.Pow(1024, power)), nil
}
