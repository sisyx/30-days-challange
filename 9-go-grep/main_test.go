package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFindInFile(t *testing.T) {
	file := filepath.Join(t.TempDir(), "input.log")
	if err := os.WriteFile(file, []byte("GET first\nno match\nGET twice GET\nlast GET"), 0600); err != nil {
		t.Fatal(err)
	}

	lines, exists := findInFile(file, "GET")
	want := []string{"GET first", "GET twice GET", "last GET"}
	if !exists || !reflect.DeepEqual(lines, want) {
		t.Fatalf("findInFile(GET) = %q, %v; want %q, true", lines, exists, want)
	}

	for _, phrase := range []string{"missing", ""} {
		lines, exists := findInFile(file, phrase)
		if exists || len(lines) != 0 {
			t.Fatalf("findInFile(%q) = %q, %v; want no matches", phrase, lines, exists)
		}
	}
}
