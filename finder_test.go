package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestSuccessUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Go Go")
	}))
	defer server.Close()

	task := NewTask(server.URL, "Go")
	task.run()

	if task.countMatch != 3 {
		t.Errorf("Count incorrect: %d != 3", task.countMatch)
	}
}
