package echo

import (
	"strings"
)

func EchoLoudly(input string) string {
	return strings.ToUpper(input)
}

func EchoRepeat(input string) string {
	return input
}
