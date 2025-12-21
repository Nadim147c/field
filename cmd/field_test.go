package cmd

import (
	"reflect"
	"testing"
)

func TestFieldN(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimiter string
		limit     int
		want      []string
	}{
		{
			name:      "basic unlimited split on space",
			input:     "a     b c",
			delimiter: " ",
			limit:     -1,
			want:      []string{"a", "b", "c"},
		},
		{
			name:      "split with limit=2 keeps remainder intact",
			input:     "a     b c",
			delimiter: " ",
			limit:     2,
			want:      []string{"a", "b c"},
		},
		{
			name:      "trailing delimiters are ignored",
			input:     "a     b c     ",
			delimiter: " ",
			limit:     -1,
			want:      []string{"a", "b", "c"},
		},
		{
			name:      "multi-char delimiter",
			input:     "a      b c",
			delimiter: "  ",
			limit:     -1,
			want:      []string{"a", "b c"},
		},
		{
			name:      "limit preserves spaces in last field",
			input:     "a v c d     c",
			delimiter: " ",
			limit:     4,
			want:      []string{"a", "v", "c", "d     c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FieldN([]byte(tt.input), tt.delimiter, tt.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FieldN(%q, %q, %d) = %q, want %q",
					tt.input, tt.delimiter, tt.limit, got, tt.want)
			}
		})
	}
}
