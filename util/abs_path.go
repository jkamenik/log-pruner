package util

import (
	"os"
	"path"
)

// AbsPath returns a full AbsPath
func AbsPath(pathName string) (string, error) {
	wd, err := os.Getwd()
	return absPath(pathName, wd), err
}

// Either return path or base joined with path
func absPath(pathName, base string) string {
	if path.IsAbs(pathName) {
		return pathName
	}

	return path.Join(base, pathName)
}
