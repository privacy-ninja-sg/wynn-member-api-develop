package utils

import (
	"strconv"
)

// StrToInt : function for convert string to integer
func StrToInt(s string) (int, error) {
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return res, nil
}

// IntToStr : function for convert integer to string
func IntToStr(i int) string {
	res := strconv.Itoa(i)
	return res
}
