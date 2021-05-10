package main

import "strings"

// match returns true if src matches pattern,
// false otherwise.
func match(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}
