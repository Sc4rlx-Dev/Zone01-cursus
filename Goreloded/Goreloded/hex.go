package main

import (
	"strconv"
)

func hexToInt(hex string) string {
	val, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return hex
	}
	return strconv.FormatInt(val, 10)
}
