package scanner_test

import (
	"testing"
)

func TestNew(t *testing.T) {
	scan := initFs(t, "/logs/testing")

	if scan.Name() != "test.log" {
		t.Errorf("file name '%s' is not expected", scan.Name())
	}

	if scan.AbsPath() != "/logs/testing/test.log" {
		t.Errorf("File's abs path is not set correctly (%s)", scan.AbsPath())
	}

	if scan.IsProtected() {
		t.Error("File should not be protected.")
	}
}

func TestProtectedFiles(t *testing.T) {
	scan := initFs(t, "/logs/testing")

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
