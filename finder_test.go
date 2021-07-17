package main

import (
	"testing"
)

func TestCountSubStr(t *testing.T) {
	cases := []struct {
		str         string
		searchWord  string
		cntFindWord int
	}{
		{
			"Go123123123",
			"Go",
			1,
		},
		{
			"GO GoGoGo",
			"Go",
			3,
		},
		{
			"Go Go Go Go Go Go Go Go",
			"Go",
			8,
		},
	}
	for _, c := range cases {
		task := Task{
			searchWord: c.searchWord,
		}
		cnt := task.findMatch([]byte(c.str))
		if cnt != c.cntFindWord {
			t.Errorf("task.findMatch(%q, %q) == %d, cntFindWord %d", c.str, c.searchWord, cnt, c.cntFindWord)
		}
	}
}
