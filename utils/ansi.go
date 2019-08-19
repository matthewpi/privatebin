package utils

import (
	"regexp"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var regex = regexp.MustCompile(ansi)

// StripANSI strips ansi out of a string.
func StripANSI(str string) string {
	return regex.ReplaceAllString(str, "")
}
