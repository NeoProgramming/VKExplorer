package core

import (
	"math"
	"strconv"
	"time"
)

func Atoi(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func Ttoa(t time.Time) string {
	if t == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		return "---"
	}
	return t.Format("06-01-02 15:04")
}

func Tmtoa(t int64) string {
	if t == 0 {
		return "---"
	}
	return time.Unix(t, 0).Format("06-01-02 15:04")
}

/*
func minTime(t ...time.Time) time.Time {
	mt := time.Unix(1<<63-62135596801, 999999999)
	for _, t := range t {
		if t.Before(mt) {
			mt = t
		}
	}
	return mt
}

func maxTime(t ...time.Time) time.Time {
	mt := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, t := range t {
		if mt.Before(t) {
			mt = t
		}
	}
	return mt
}*/
func minTime(t ...int64) int64 {
	var mt int64 = math.MaxInt64
	for _, t := range t {
		if t < mt {
			mt = t
		}
	}
	return mt
}

func maxTime(t ...int64) int64 {
	var mt int64 = 0
	for _, t := range t {
		if mt < t {
			mt = t
		}
	}
	return mt
}
