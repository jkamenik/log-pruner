package scanner

import (
	"strings"
)

type sortFile []File

func (a sortFile) Len() int      { return len(a) }
func (a sortFile) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortFile) Less(i, j int) bool {
	iTime := a[i].ModTime()
	jTime := a[j].ModTime()

	// sort on time and if they are equal on name.
	if iTime.Equal(jTime) {
		return strings.Compare(a[i].AbsPath(), a[j].AbsPath()) < 0
	}

	return a[i].ModTime().Before(a[j].ModTime())
}
