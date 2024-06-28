package older_test

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/ezebunandu/older"

	"github.com/google/go-cmp/cmp"
)

func TestFilesOlderThan25DaysCorrectlyListFilesInMapFS(t *testing.T) {
	t.Parallel()
	now := time.Now()
	fsys := fstest.MapFS{
		"file.go":                {ModTime: now},
		"subfolder/subfolder.go": {ModTime: now.Add(-time.Minute)},
		"subfolder2/another.go":  {ModTime: now},
		"subfolder2/file.go":     {ModTime: now.Add(-time.Minute)},
	}
	want := []string{
		"subfolder/subfolder.go",
		"subfolder/subfolder.go",
	}
	got := older.Files(fsys, 10*time.Second)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
