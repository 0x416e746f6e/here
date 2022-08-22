// Package here ...
package here

import "strings"

// Doc returns trimmed and unindented string.
//
// Pre-requisites are as follows:
//
//   - Raw input must start and end with a new-line `\n`.
//   - Only tab characters `\t` are de-dented.
func Doc(s string) string {
	if len(s) == 0 {
		return s
	}
	if s[0] != '\n' {
		return s
	}
	tail := len(s) - 1
	for ; tail >= 0; tail-- {
		if s[tail] == '\t' {
			continue
		}
		if s[tail] == '\n' {
			break
		}
		return s
	}

	s = s[1:tail]
	ls := strings.Split(s, "\n")

	indent := indent(ls)
	for i, l := range ls {
		if len(l) > 0 {
			ls[i] = l[indent:]
		}
	}

	return strings.Join(ls, "\n")
}

// countLeadingTabs returns the count of consecutive leading tab characters for
// the string.
func countLeadingTabs(l string) int {
	for i, c := range l {
		if c != '\t' {
			return i
		}
	}
	return len(l)
}

// indent returns min count of leading tabs for the slice of strings.
func indent(ls []string) int {
	indent := countLeadingTabs(ls[0])

	for _, l := range ls {
		if len(l) > 0 {
			tabs := countLeadingTabs(l)
			if tabs < indent {
				indent = tabs
			}
		}
	}

	return indent
}
