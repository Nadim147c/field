package cmd

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Range is data type the indicates a range of item for a string slice
type Range struct {
	Reversed   bool
	Exact      bool
	Start, End int
}

// ParseRange parses a string into a Range
func ParseRange(str string, reversed bool) (*Range, error) {
	if str == "" {
		return nil, errors.New("empty range string")
	}

	if str == ":" {
		return &Range{Reversed: reversed, Start: 0, End: math.MaxInt}, nil
	}

	// Check for exact values (no operators)
	if !strings.ContainsRune(str, ':') {
		start, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("failed to parse range: %v", err)
		}
		return &Range{Exact: true, Start: start, Reversed: reversed}, nil
	}

	parts := strings.Split(str, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid range string: %q", str)
	}

	start := 0
	end := math.MaxInt

	if parts[0] != "" {
		n, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse range: %v", err)
		}
		start = n
	}

	if parts[1] != "" {
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse range: %v", err)
		}
		end = n
	}

	return &Range{Start: start, End: end, Reversed: reversed}, nil
}

// Select selects item of a []string according to the bound
func (r *Range) Select(s []string) []string {
	length := len(s)
	if length == 0 {
		return nil
	}

	// Handle exact selection
	if r.Exact {
		idx := r.Start
		if idx < 0 {
			idx = length + idx
		} else if idx > 0 {
			idx = idx - 1 // Convert 1-based to 0-based
		}

		// idx == 0 remains 0 (first element)
		if idx < 0 || idx >= length {
			return nil
		}

		return []string{s[idx]}
	}

	start, end := r.Start, r.End

	if start < 0 {
		start = length + start
	} else if start > 0 {
		start = start - 1 // Convert 1-based to 0-based
	}

	if end < 0 {
		end = length + end
	} else if end > 0 {
		end = end - 1 // Convert 1-based to 0-based
	}

	if start > end || start >= length {
		return nil
	}

	start = max(0, min(start, length-1))
	end = max(0, min(end, length-1))

	if r.Reversed {
		sliced := slices.Clone(s[start : end+1])
		slices.Reverse(sliced)
		return sliced
	}

	return s[start : end+1]
}
