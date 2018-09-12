package scanner

import (
	"os"
	"sort"
	"time"

	"github.com/spf13/afero"
)

// AppFs is the base Afero FS type.  For testing this would be mapped to memory or something else.
var AppFs = afero.NewOsFs()

// Path is the collection abstration for a group of files that can be pruned.
type Path struct {
	basePath string
	files    []File

	fileAgeout int

	totalSize   int64
	totalPruned int64
}

// NewPath creates a new Path object containing the intial list of walked files from basePath.  These files can be later scanned for size or age.
func NewPath(basePath string, fileAgeout int) *Path {
	path := &Path{basePath: basePath, fileAgeout: fileAgeout}

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
	totalSize := int64(0)
	ptrTotalSize := &totalSize
	err := afero.Walk(AppFs, path.basePath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		*ptrTotalSize = *ptrTotalSize + f.Size()

		fileList = append(fileList, NewFile(f, path))
		return nil
	})

	if err != nil {
		return nil, err
	}

	// sort
	sort.Sort(sortFile(fileList))

	path.files = fileList
	path.totalSize = totalSize
	path.totalPruned = 0

	return path.files, nil
}

// TotalSize returns the total size in bytes for the scanned files.  Useful in seeing now much will be saved when saving.
func (path *Path) TotalSize() int64 {
	return path.totalSize
}

// TotalAfterPrune returns the total size minus any files marked for pruning
func (path *Path) TotalAfterPrune() int64 {
	return path.totalSize - path.totalPruned
}

// MarkOldFiles runs the list and
func (path *Path) MarkOldFiles() {
	before := time.Now().AddDate(0, 0, -1*path.fileAgeout)

	for i, file := range path.files {
		if file.ModTime().Before(before) {
			path.files[i].Prune(true)
		}
	}

	path.rescanForPruned()
}

// rescan pruned
func (path *Path) rescanForPruned() {
	path.totalPruned = 0
	for _, file := range path.files {
		if file.WillPrune() {
			path.totalPruned = path.totalPruned + file.Size()
		}
	}
}

// TODO: SizeScan
