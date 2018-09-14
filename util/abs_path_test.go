package util

import "testing"

func TestAbs(t *testing.T) {
	examples := map[string]string{
		".":          "/foo",
		"./bin":      "/foo/bin",
		"../this":    "/this",
		"../../that": "/that",
		"/tmp":       "/tmp",
	}

	for key, value := range examples {
		actual := absPath(key, "/foo")
		if value != actual {
			t.Errorf("'%s' should have become '%s', but became '%s'", key, value, actual)
		}
	}
}
