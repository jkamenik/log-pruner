package scanner

import "github.com/spf13/afero"

// AppFs is the base Afero FS type.  For testing this would be mapped to memory or something else.
var AppFs = afero.NewOsFs()
