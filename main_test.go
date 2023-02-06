package main

import "testing"

func TestChuckNorris(t *testing.T) {
	tables := []struct {
		input string
		want  string
	}{
		{"C", "0 0 00 0000 0 00"},
		{"CC", "0 0 00 0000 0 000 00 0000 0 00"},
		{"Hi <3", "0 0 00 00 0 0 00 000 0 00 00 0 0 0 00 00 0 0 00 0 0 0 00 000000 0 0000 00 000 0 00 00 00 0 00"},
	}

	for _, table := range tables {
		got := ChuckNorris(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \nwant: %v, \n got: %v", table.want, got)
		}
	}
}
