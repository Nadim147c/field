package cmd

import (
	"bytes"
	"unicode/utf8"
)

const minResultSize int = 30

// FieldN takes a string and a delimiter and returns n numbers of fields
// seperated by the delimiter. If n is negative the returns all fields.
func FieldN(s []byte, delimiter string, n int) []string {
	if len(s) == 0 || delimiter == "" {
		return nil
	}
	if n == 1 {
		return []string{string(s)}
	}

	size := len(s)
	d := len(delimiter)
	result := make([]string, 0, minResultSize)

	start := 0
	found := 0

	delim := []byte(delimiter)

	for i := 0; i+d <= size; {
		if bytes.HasPrefix(s[i:], delim) {
			if start < i {
				result = append(result, string(s[start:i]))
				found++
			}
			i += d
			for i+d <= size && bytes.HasPrefix(s[i:], delim) {
				i += d
			}

			if n > 0 && found+1 == n {
				result = append(result, string(s[i:]))
				return result
			}

			start = i
		} else {
			i++
		}
	}

	if start < size {
		result = append(result, string(s[start:]))
	}

	return result
}

// Pred is a function that returns true if the rune should be treated as a
// separator.
type Pred func(rune) bool

// FieldNPred splits s into at most n fields, separated by runes
// where pred(r) == true. If n < 0, it returns all fields.
// Consecutive separators are treated as one (like strings.Fields).
func FieldNPred(s []byte, pred Pred, n int) []string {
	if len(s) == 0 || delimiter == "" {
		return nil
	}
	if n == 1 {
		return []string{string(s)}
	}

	result := make([]string, 0, minResultSize)
	start := 0
	found := 0

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if pred(r) {
			if start < i {
				result = append(result, string(s[start:i]))
				found++
			}
			i += size
			// skip consecutive separators
			for i < len(s) {
				r2, sz2 := utf8.DecodeRune(s[i:])
				if !pred(r2) {
					break
				}
				i += sz2
			}

			if n > 0 && found+1 == n {
				result = append(result, string(s[i:]))
				return result
			}
			start = i
		} else {
			i += size
		}
	}

	if start < len(s) {
		result = append(result, string(s[start:]))
	}

	return result
}
