package scanner

import "os"
import "errors"

// ErrFileProtected defined attempts modify protected files
var ErrFileProtected = errors.New("File is protected")

// File wraps the os.FileInfo interface and
type File struct {
	os.FileInfo

	absPath string
	prune   bool
	protect bool
}

// NewFile upconverts a os.FileInfo object into a prunable file
func NewFile(fileInfo os.FileInfo, absPath string) File {
	return File{fileInfo, absPath, false, false}
}

// AbsPath returns the full ABS path to the file
func (file *File) AbsPath() string {
	return file.absPath
}

// WillPrune returns true of the file is marked for pruning
func (file *File) WillPrune() bool {
	if file.protect {
		return false
	}
	return file.prune
}

// IsProtected returns true if the file is protected
func (file *File) IsProtected() bool {
	return file.protect
}

// Prune marks a file as prunable or not.  Throws an error if the file is protected
func (file *File) Prune(setPrune bool) error {
	if file.protect {
		return ErrFileProtected
	}
	file.prune = setPrune
	return nil
}

// Protect marks a file preventing it from being pruned
func (file *File) Protect(setProtect bool) {
	file.protect = setProtect
}
