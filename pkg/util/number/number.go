package number

import (
	"strconv"
	"strings"
	"unicode"
)

const (
	kbFactor      = 1024
	mbFactor      = 1024 * kbFactor
	gbFactor      = 1024 * mbFactor
	minUnitLength = 2
)

func ParseFloat(value string) float64 {
	res, _ := strconv.ParseFloat(strings.TrimFunc(value, func(r rune) bool {
		return !unicode.IsNumber(r)
	}), 64)

	return res
}

func ParseBytes(value string) float64 {
	if value == "" {
		return 0
	}

	res, _ := strconv.ParseFloat(strings.TrimFunc(value, func(r rune) bool {
		return !unicode.IsNumber(r)
	}), 64)

	if len(value) < minUnitLength {
		return res
	}

	unit := strings.ToLower(strings.TrimSpace(value[len(value)-2:]))

	switch unit {
	case "kb":
		res *= kbFactor
	case "mb":
		res *= mbFactor
	case "gb":
		res *= gbFactor
	}

	return res
}
