// https://github.com/google/codesearch
// https://golang.org/pkg/bufio/#NewScanner

package search

import (
	"testing"
)

func equal(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
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

func TestSearch(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{
			in:   "#Hauptspeise",
			want: []string{"Kürbis-Pesto.txt", "Basilikum-Pesto.txt", "Apfel-Mangold-Tarte.txt"},
		},
	}

	for _, c := range cases {
		got, err := Search("data/", c.in)
		if err != nil {
			t.Error("Function returned an error.")
		}
		if equal(got, c.want) {
			t.Errorf("Search(\"data/\", %v) == %v, want %v.\n", c.in, got, c.want)
		}
	}
}

func TestSearchResult(t *testing.T) {
	cases := []struct {
		in   *SearchResult
		want []string
	}{
		{
			in:   &SearchResult{Dir: "data/", Tag: "#Hauptspeise"},
			want: []string{"Kürbis-Pesto.txt", "Basilikum-Pesto.txt", "Apfel-Mangold-Tarte.txt"},
		},
	}

	for _, c := range cases {
		err := c.in.Search()
		if err != nil {
			t.Error("Function returned an error.")
		}
		if equal(c.in.Files, c.want) {
			t.Errorf("%+v.Search() == %v, want %v.", c.in, c.in.Files, c.want)
		}
	}
}
