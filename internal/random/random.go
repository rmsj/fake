package random

import (
	"math/rand"
	"time"
)

// StringFromSlice returns a random element from the given slice of string
func StringFromSlice(s []string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s[0]
	}
	rand.Seed(time.Now().UnixNano())
	return s[rand.Intn(len(s)-1)]
}
