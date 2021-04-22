// Package xterm used to set or strip colors in terminal.
package xterm

import "regexp"

// stripAnsiRegexp is a regular expression to clean ANSI Control sequences
// feat https://github.com/chalk/ansi-regex
var stripAnsiRegexp = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")

// Strip any ANSI escape codes from string.
func Strip(s string) string {
	return stripAnsiRegexp.ReplaceAllString(s, "")
}
