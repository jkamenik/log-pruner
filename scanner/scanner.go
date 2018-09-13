package scanner

import (
	"log"
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

	fileAgeoutDays int
	maxSizeBytes   int64

	totalSize   int64
	totalPruned int64
}

// NewPath creates a new Path object containing the intial list of walked files from basePath.  These files can be later scanned for size or age.
func NewPath(basePath string, fileAgeoutDays int, maxSizeBytes int64) *Path {
	log.Printf("NewPath for %s, %dDays, %dBytes", basePath, fileAgeoutDays, maxSizeBytes)
	path := &Path{basePath: basePath, fileAgeoutDays: fileAgeoutDays, maxSizeBytes: maxSizeBytes}

	path.ReadDir()

	return path
}

// Files returns the full list of walked files
func (path *Path) Files() []File {
	return path.files
}

// ReadDir is a destructive full re-read of the directory.  All existing, prune details are lost.
func (path *Path) ReadDir() ([]File, error) {
	log.Print("ReadDir()")
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

		log.Printf("%s (%dBytes) %s", path, f.Size(), f.ModTime())

		*ptrTotalSize = *ptrTotalSize + f.Size()

		fileList = append(fileList, NewFile(f, path))
		return nil
	})

	if err != nil {
		return nil, err
	}

	// sort
	log.Print("Sorting files")
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
	log.Printf("TotalAfterPrune: %dbytes, %dbytes prunable", path.totalSize, path.totalPruned)
	return path.totalSize - path.totalPruned
}

// MarkOldFiles runs the list and
func (path *Path) MarkOldFiles() {
	log.Printf("MarkOldFiles older then %ddays", path.fileAgeoutDays)
	before := time.Now().AddDate(0, 0, -1*path.fileAgeoutDays)

	for i, file := range path.files {
		if file.ModTime().Before(before) {
			log.Printf("%s older then 3 days, marked for pruning", file.AbsPath())
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
func (path *Path) MarkFileUntilFit() {
	log.Printf("MarkFileUntilFit into %dbytes, currently %dbytes", path.maxSizeBytes, path.TotalAfterPrune())
	for i, file := range path.files {
		if path.TotalAfterPrune() <= path.maxSizeBytes {
			log.Printf("Size has been reached, stopping.")
			break
		}
		if file.WillPrune() {
			// skip files already pruned
			continue
		}

		log.Printf("%s will be pruned, saving %dbytes", file.AbsPath(), file.Size())
		if err := path.files[i].Prune(true); err == nil {
			path.totalPruned = path.totalPruned + file.Size()
		} else {
			log.Print(err)
		}
	}

	// just for run rescan everything
	path.rescanForPruned()
}

// Prune will remove files that were marked for removing.
func (path *Path) Prune() error {
	for _, file := range path.files {
		if file.WillPrune() {
			log.Printf("Pruning %s", file.AbsPath())
			err := AppFs.Remove(file.AbsPath())
			if err != nil {
				log.Print(err)
				return err
			}
		}
	}

	_, err := path.ReadDir()
	return err
}
