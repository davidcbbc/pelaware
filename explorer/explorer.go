package explorer

import (
	"io/fs"
	"path/filepath"
)

func MapFiles(dir string) []string {
	var files []string

	error := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)

		return nil
	})

	if error != nil {
		panic(error)
	}

	return files
}
