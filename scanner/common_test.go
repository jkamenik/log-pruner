package scanner_test

import (
	"log"
	"path"
	"testing"

	"github.com/jkamenik/log-pruner/scanner"
	"github.com/spf13/afero"
)

var AppFs = afero.NewMemMapFs()

func init() {
	scanner.AppFs = AppFs
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)
}

// Init a memory logging fs with a single test file.  Returns the file
func initFs(t *testing.T, basePath string) scanner.File {
	testFile := path.Join(basePath, "test.log")
	AppFs.RemoveAll(basePath)
	AppFs.MkdirAll(basePath, 0644)
	afero.WriteFile(AppFs, testFile, []byte("Some test log"), 0644)
	file, err := AppFs.Stat(testFile)
	if err != nil {
		t.Fatal(err)
	}

	return scanner.NewFile(file, testFile)
}
