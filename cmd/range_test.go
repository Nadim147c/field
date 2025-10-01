package cmd

import (
	"testing"
)

func TestRange_Select(t *testing.T) {
	testSlice := []string{"a", "b", "c", "d", "e"}

	tests := []struct {
		name     string
		rangeStr string
		reversed bool
		input    []string
		want     []string
	}{
		// Exact selection tests (reversed doesn't affect single elements)
		{
			name:     "exact positive index first element",
			rangeStr: "1",
			reversed: false,
			input:    testSlice,
			want:     []string{"a"},
		},
		{
			name:     "exact positive index middle element",
			rangeStr: "3",
			reversed: false,
			input:    testSlice,
			want:     []string{"c"},
		},
		{
			name:     "exact positive index last element",
			rangeStr: "5",
			reversed: false,
			input:    testSlice,
			want:     []string{"e"},
		},
		{
			name:     "exact negative index last element",
			rangeStr: "-1",
			reversed: false,
			input:    testSlice,
			want:     []string{"e"},
		},
		{
			name:     "exact negative index second last element",
			rangeStr: "-2",
			reversed: false,
			input:    testSlice,
			want:     []string{"d"},
		},
		{
			name:     "exact index 0",
			rangeStr: "0",
			reversed: false,
			input:    testSlice,
			want:     []string{"a"},
		},
		{
			name:     "exact index out of bounds positive",
			rangeStr: "10",
			reversed: false,
			input:    testSlice,
			want:     nil,
		},
		{
			name:     "exact index out of bounds negative",
			rangeStr: "-10",
			reversed: false,
			input:    testSlice,
			want:     nil,
		},

		// Range selection tests - not reversed
		{
			name:     "range from start to end",
			rangeStr: "1:5",
			reversed: false,
			input:    testSlice,
			want:     []string{"a", "b", "c", "d", "e"},
		},
		{
			name:     "range partial from start",
			rangeStr: "1:3",
			reversed: false,
			input:    testSlice,
			want:     []string{"a", "b", "c"},
		},
		{
			name:     "range partial from middle",
			rangeStr: "2:4",
			reversed: false,
			input:    testSlice,
			want:     []string{"b", "c", "d"},
		},
		{
			name:     "range partial to end",
			rangeStr: "3:5",
			reversed: false,
			input:    testSlice,
			want:     []string{"c", "d", "e"},
		},
		{
			name:     "range with negative end",
			rangeStr: "2:-1",
			reversed: false,
			input:    testSlice,
			want:     []string{"b", "c", "d", "e"},
		},
		{
			name:     "range with negative start and end",
			rangeStr: "-4:-2",
			reversed: false,
			input:    testSlice,
			want:     []string{"b", "c", "d"},
		},
		{
			name:     "range with empty start",
			rangeStr: ":3",
			reversed: false,
			input:    testSlice,
			want:     []string{"a", "b", "c"},
		},
		{
			name:     "range with empty end",
			rangeStr: "3:",
			reversed: false,
			input:    testSlice,
			want:     []string{"c", "d", "e"},
		},
		{
			name:     "range with both empty",
			rangeStr: ":",
			reversed: false,
			input:    testSlice,
			want:     []string{"a", "b", "c", "d", "e"},
		},

		// Range selection tests - reversed
		{
			name:     "reversed range from start to end",
			rangeStr: "1:5",
			reversed: true,
			input:    testSlice,
			want:     []string{"e", "d", "c", "b", "a"},
		},
		{
			name:     "reversed range partial from start",
			rangeStr: "1:3",
			reversed: true,
			input:    testSlice,
			want:     []string{"c", "b", "a"},
		},
		{
			name:     "reversed range partial from middle",
			rangeStr: "2:4",
			reversed: true,
			input:    testSlice,
			want:     []string{"d", "c", "b"},
		},
		{
			name:     "reversed range partial to end",
			rangeStr: "3:5",
			reversed: true,
			input:    testSlice,
			want:     []string{"e", "d", "c"},
		},
		{
			name:     "reversed range with negative end",
			rangeStr: "2:-2",
			reversed: true,
			input:    testSlice,
			want:     []string{"d", "c", "b"},
		},
		{
			name:     "reversed range with negative start and end",
			rangeStr: "-4:-2",
			reversed: true,
			input:    testSlice,
			want:     []string{"d", "c", "b"},
		},
		{
			name:     "reversed range with empty start",
			rangeStr: ":3",
			reversed: true,
			input:    testSlice,
			want:     []string{"c", "b", "a"},
		},
		{
			name:     "reversed range with empty end",
			rangeStr: "3:",
			reversed: true,
			input:    testSlice,
			want:     []string{"e", "d", "c"},
		},
		{
			name:     "reversed range with both empty",
			rangeStr: ":",
			reversed: true,
			input:    testSlice,
			want:     []string{"e", "d", "c", "b", "a"},
		},

		// Edge cases
		{
			name:     "empty slice",
			rangeStr: "1:3",
			reversed: false,
			input:    []string{},
			want:     nil,
		},
		{
			name:     "single element slice exact",
			rangeStr: "1",
			reversed: false,
			input:    []string{"x"},
			want:     []string{"x"},
		},
		{
			name:     "single element slice exact reversed",
			rangeStr: "1",
			reversed: true,
			input:    []string{"x"},
			want:     []string{"x"}, // Single element reversed is still the same
		},
		{
			name:     "two element slice range",
			rangeStr: "1:2",
			reversed: false,
			input:    []string{"x", "y"},
			want:     []string{"x", "y"},
		},
		{
			name:     "two element slice range reversed",
			rangeStr: "1:2",
			reversed: true,
			input:    []string{"x", "y"},
			want:     []string{"y", "x"},
		},
		{
			name:     "range start greater than end",
			rangeStr: "4:2",
			reversed: false,
			input:    testSlice,
			want:     nil,
		},
		{
			name:     "range start greater than end reversed",
			rangeStr: "4:2",
			reversed: true,
			input:    testSlice,
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := ParseRange(tt.rangeStr, tt.reversed)
			if err != nil {
				t.Fatalf("ParseRange(%q, %v) failed: %v", tt.rangeStr, tt.reversed, err)
			}

			got := r.Select(tt.input)

			if !equalSlices(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_Select_BoundaryConditions(t *testing.T) {
	tests := []struct {
		name     string
		rangeStr string
		reversed bool
		input    []string
		want     []string
	}{
		{
			name:     "range at upper boundary",
			rangeStr: "4:5",
			reversed: false,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"d", "e"},
		},
		{
			name:     "range at upper boundary reversed",
			rangeStr: "4:5",
			reversed: true,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"e", "d"},
		},
		{
			name:     "range beyond upper boundary clamped",
			rangeStr: "4:10",
			reversed: false,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"d", "e"},
		},
		{
			name:     "range beyond upper boundary clamped reversed",
			rangeStr: "4:10",
			reversed: true,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"e", "d"},
		},
		{
			name:     "range beyond lower boundary clamped",
			rangeStr: "-10:2",
			reversed: false,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"a", "b"},
		},
		{
			name:     "range beyond lower boundary clamped reversed",
			rangeStr: "-10:2",
			reversed: true,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"b", "a"},
		},
		{
			name:     "zero index range start",
			rangeStr: "0:2",
			reversed: false,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"a", "b"},
		},
		{
			name:     "zero index range start reversed",
			rangeStr: "0:2",
			reversed: true,
			input:    []string{"a", "b", "c", "d", "e"},
			want:     []string{"b", "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := ParseRange(tt.rangeStr, tt.reversed)
			if err != nil {
				t.Fatalf("ParseRange(%q) failed: %v", tt.rangeStr, err)
			}

			got := r.Select(tt.input)

			if !equalSlices(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange_Select_InvalidRange(t *testing.T) {
	testSlice := []string{"a", "b", "c", "d", "e"}

	// Test that invalid ranges return nil (both normal and reversed)
	invalidRanges := []struct {
		rangeStr string
		reversed bool
	}{
		{"10:5", false},
		{"10:5", true},
		{"3:1", false},
		{"3:1", true},
		{"-1:-3", false},
		{"-1:-3", true},
		{"100:200", false},
		{"100:200", true},
	}

	for _, tt := range invalidRanges {
		t.Run(tt.rangeStr, func(t *testing.T) {
			r, err := ParseRange(tt.rangeStr, tt.reversed)
			if err != nil {
				t.Fatalf("ParseRange(%q) failed: %v", tt.rangeStr, err)
			}

			got := r.Select(testSlice)
			if got != nil {
				t.Errorf("Select() for invalid range %q (reversed=%v) = %v, want nil", tt.rangeStr, tt.reversed, got)
			}
		})
	}
}

// Helper function to compare slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
