// Package tests contains supporting code for running tests.
package tests

import (
	"fmt"
	"regexp"

	"runtime"
)

// Success and failure markers.
var (
	success = "\u2713"
	failed  = "\u2717"
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	purple  = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[37m"
	white   = "\033[97m"
)

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}

// Success prints in green the given string
func Success(before string, s string) string {
	return before + fmt.Sprintf("%s%s\t%s%s", green, success, s, reset)
}

// Failed prints in red the given string
func Failed(before string, s string) string {
	return before + fmt.Sprintf("%s%s\\%s%s", red, failed, s, reset)
}

// Reset returns the code to reset color
func Reset() string {
	return reset
}

// Red returns the code to red color
func Red() string {
	return red
}

// Green returns the code to green color
func Green() string {
	return green
}

// Yellow returns the code to yellow color
func Yellow() string {
	return yellow
}

// Blue returns the code to blue color
func Blue() string {
	return blue
}

// Purple returns the code to purple color
func Purple() string {
	return purple
}

// Cyan returns the code to cyan color
func Cyan() string {
	return cyan
}

// Gray returns the code to gray color
func Gray() string {
	return gray
}

// White returns the code to white color
func White() string {
	return white
}

// ValidEmail checks if a string is a valid email address
func ValidEmail(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
