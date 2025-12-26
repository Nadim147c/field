package cmd

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func FuzzFieldN(f *testing.F) {
	// Seed corpus
	f.Add([]byte("a,b,c"), ",", 1)
	f.Add([]byte("a,,b,,c"), ",", -1)
	f.Add([]byte("a--b--c"), "--", 2)
	f.Add([]byte(""), ",", -1)
	f.Add([]byte("no delimiter"), "|", -1)

	f.Fuzz(func(t *testing.T, s []byte, delim string, n int) {
		// Avoid meaningless cases
		if delim == "" {
			return
		}

		res := FieldN(s, delim, n)

		// n == 1 invariant
		if n == 1 && len(s) > 0 {
			if len(res) != 1 || res[0] != string(s) {
				t.Fatalf("n=1 violated: %q vs %q", res, s)
			}
		}
	})
}

func FuzzFieldNFunc(f *testing.F) {
	// Stable predicates
	isSpace := func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n'
	}

	f.Add([]byte("a b  c"), 0)
	f.Add([]byte("a,,b,,,c"), -1)
	f.Add([]byte("日本 語 テ ス ト"), -1)

	f.Fuzz(func(t *testing.T, s []byte, n int) {
		// Force valid UTF-8 sometimes
		if !utf8.Valid(s) {
			return
		}

		res := FieldNFunc(s, isSpace, n)

		// n == 1 invariant
		if n == 1 && len(s) > 0 {
			if len(res) != 1 || res[0] != string(s) {
				t.Fatalf("n=1 violated: %q vs %q", res, s)
			}
		}

		// No empty fields (like strings.Fields)
		for i, field := range res {
			if field == "" {
				t.Fatalf("empty field at %d: %#v", i, res)
			}
		}
	})
}

func FuzzFieldNFunc_Stdlib(f *testing.F) {
	isSpace := func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n'
	}

	f.Fuzz(func(t *testing.T, s []byte) {
		if !utf8.Valid(s) {
			return
		}

		got := FieldNFunc(s, isSpace, -1)
		want := strings.FieldsFunc(string(s), isSpace)

		if len(got) != len(want) {
			t.Fatalf("len mismatch: %v vs %v", got, want)
		}

		for i := range got {
			if got[i] != want[i] {
				t.Fatalf("mismatch at %d: %q vs %q", i, got[i], want[i])
			}
		}
	})
}
