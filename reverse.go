// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	front := 0
	back := len(r) - 1
	for front < len(r)/2 {
		r[front], r[back] = r[back], r[front]
		front++
		back--
	}
	return string(r)
}
