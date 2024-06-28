package findFiles_test

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/ezebunandu/findFiles"

	"github.com/google/go-cmp/cmp"
)

func TestFilesOlderThan25DaysCorrectlyListFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                &fstest.MapFile{ModTime: time.Date(2024, 5, 28, 12, 0, 0, 0, time.UTC)},
		"subfolder/subfolder.go": &fstest.MapFile{ModTime: time.Date(2024, 6, 2, 12, 0, 0, 0, time.UTC)},
		"subfolder2/another.go":  &fstest.MapFile{ModTime: time.Date(2024, 6, 23, 12, 0, 0, 0, time.UTC)},
		"subfolder2/file.go":     &fstest.MapFile{ModTime: time.Date(2024, 6, 27, 12, 0, 0, 0, time.UTC)},
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
	}
	got := findFiles.FilesOlderThan25Days(fsys)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
