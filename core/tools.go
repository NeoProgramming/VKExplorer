package core

import (
	"strconv"
	"time"
)

func Atoi(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func minTime(t ...time.Time) time.Time {
	mt := time.Unix(9223372036854775807, 999999999)
	for _, t := range t {
		if mt.Before(t) {
			mt = t
		}
	}
	return mt
}

func maxTime(t ...time.Time) time.Time {
	mt := time.Unix(0, 0)
	for _, t := range t {
		if t.Before(mt) {
			mt = t
		}
	}
	return mt
}
