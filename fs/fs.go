package fs

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func listFiles(fsys fs.FS, dir string) ([]string, error) {
	var fileList []string

	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %q: %w", dir, err)
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			subDirFiles, err := listFiles(fsys, path)
			if err != nil {
				return nil, err
			}
			fileList = append(fileList, subDirFiles...)
		} else {
			fileList = append(fileList, path)
		}
	}

	return fileList, nil
}

func Walk(fsys fs.FS) ([]string, error) {
	return listFiles(fsys, ".")
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func MkDirIfNone(dir string) error {
	if DirExists(dir) {
		return nil
	}
	return os.Mkdir(dir, 0o755)
}
