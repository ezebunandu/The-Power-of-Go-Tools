package older

import (
	"io/fs"
	"time"
)

func Files(fsys fs.FS, age time.Duration) (paths []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		return nil
	})
	return paths
}
