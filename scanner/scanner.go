package scanner

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/afero"
)

// AppFs is the base Afero FS type.  For testing this would be mapped to memory or something else.
var AppFs = afero.NewOsFs()

// Path is the collection abstration for a group of files that can be pruned.
type Path struct {
	basePath string
	files    []File

	// totalSize
	// total
}

// NewPath creates a new Path object containing the intial list of walked files from basePath.  These files can be later scanned for size or age.
func NewPath(basePath string) *Path {
	path := &Path{basePath: basePath}

	path.ReadDir()

	return path
}

// Files returns the full list of walked files
func (path *Path) Files() []File {
	return path.files
}

// ReadDir is a destructive full re-read of the directory.  All existing, prune details are lost.
func (path *Path) ReadDir() ([]File, error) {
	path.files = nil

	fileList := []File{}
	err := afero.Walk(AppFs, path.basePath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		fmt.Println(path)

		fileList = append(fileList, NewFile(f, path))
		return nil
	})

	if err != nil {
		return nil, err
	}

	// sort
	sort.Sort(sortFile(fileList))

	path.files = fileList

	return path.files, nil
}

// MarkOldFiles runs the list and
func (path *Path) MarkOldFiles() {

}

// TODO: AgeScan
// TODO: SizeScan
