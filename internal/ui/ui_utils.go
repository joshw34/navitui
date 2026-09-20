package ui

import (
	"strconv"
)

func pluralize(s string, n int) string {
	if n > 1 {
		s += "s"
	}
	return s
}

func yearLabel(y int) string {
	if y == 0 {
		return "(unknown)"
	}
	return strconv.Itoa(y)
}
