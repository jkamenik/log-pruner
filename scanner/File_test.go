package scanner_test

import (
	"os"
	"path"
	"testing"

	"github.com/jkamenik/log-pruner/scanner"
	"github.com/spf13/afero"
)

var AppFs = afero.NewMemMapFs()

func init() {
	scanner.AppFs = AppFs
}

// Init a memory logging fs with a single test file.  Returns the file
func initFs(t *testing.T, basePath string) os.FileInfo {
	testFile := path.Join(basePath, "test.log")
	AppFs.MkdirAll(basePath, 0644)
	afero.WriteFile(AppFs, testFile, []byte("Some test log"), 0644)
	file, err := AppFs.Stat(testFile)
	if err != nil {
		t.Fatal(err)
	}

	return file
}

func TestNew(t *testing.T) {
	file := initFs(t, "/logs/testing")
	scan := scanner.New(file)

	if scan.Name() != "test.log" {
		t.Errorf("file name '%s' is not expected", scan.Name())
	}

	if scan.IsProtected() {
		t.Error("File should not be protected.")
	}
}

func TestProtectedFiles(t *testing.T) {
	file := initFs(t, "/logs/testing")
	scan := scanner.New(file)

	// if a file is set to Prune, then protected it is not pruneable
	scan.Prune(true)
	if !scan.WillPrune() {
		t.Error("File should be pruned, but isn't")
	}
	scan.Protect(true)
	if !scan.IsProtected() {
		t.Error("File should be protected, but isn't")
	}
	if scan.WillPrune() {
		t.Error("File should no longer be prunable, but is")
	}

	// a file protected cannot have prunability modifed
	if err := scan.Prune(true); err == nil {
		t.Error("File should not be allowed to be pruned, but is")
	}
	if err := scan.Prune(false); err == nil {
		t.Error("File should not be allowed to be pruned, but is")
	}
}
