package findFiles

import (
	"io/fs"
	"time"
)

func FilesOlderThan25Days(fsys fs.FS) (paths []string) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		f, err := fs.Stat(fsys, p)
		if err != nil {
			return err
		}
		modTime := f.ModTime()
		now := time.Now()
		daysSince := now.Sub(modTime).Hours() / 24
		if int(daysSince) > 25 {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
