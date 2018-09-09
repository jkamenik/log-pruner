package scanner_test

import (
	"testing"
	"time"

	"github.com/jkamenik/log-pruner/scanner"
	"github.com/spf13/afero"
)

func TestSimpleScan(t *testing.T) {
	_ = initFs(t, "/logs")

	// now add a second log in a subdir
	AppFs.MkdirAll("/logs/sub", 0644)
	afero.WriteFile(AppFs, "/logs/sub/second.log", []byte("Some test log"), 0644)

	path := scanner.NewPath("/logs")
	if len(path.Files()) < 2 {
		t.Error("Path should contain at least one file.")
	}

	for _, f := range path.Files() {
		if f.IsDir() {
			t.Error("Directories should be scanned for size")
		}
	}
}

func TestResultsOrdered(t *testing.T) {
	_ = initFs(t, "/logs")
	// Files are ususally ordered by name, so make sure it is mod time
	first := "/logs/sub/z.log"
	second := "/logs/sub/a.log"
	third := "/logs/sub/b.log"
	firstTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	secondTime := time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC)

	// now add a second log in a subdir
	AppFs.MkdirAll("/logs/sub", 0644)
	afero.WriteFile(AppFs, second, []byte("Some test log"), 0644)
	afero.WriteFile(AppFs, third, []byte("Some test log"), 0644)
	afero.WriteFile(AppFs, first, []byte("Some test log"), 0644)

	// Change the file times so that first is first, and second and third are at the same time.
	AppFs.Chtimes(first, firstTime, firstTime)
	AppFs.Chtimes(second, secondTime, secondTime)
	AppFs.Chtimes(third, secondTime, secondTime)

	path := scanner.NewPath("/logs/sub")
	if len(path.Files()) != 3 {
		t.Error("Incorrect number of files found")
	}

	if path.Files()[0].AbsPath() != first ||
		path.Files()[1].AbsPath() != second ||
		path.Files()[2].AbsPath() != third {
		t.Error("Files are not ordered correctly.")
	}
}
