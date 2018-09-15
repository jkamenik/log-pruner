package util_test

import (
	"testing"

	"github.com/jkamenik/log-pruner/util"
)

func TestAbs(t *testing.T) {
	examples := map[int]int64{
		1:  1073741824,
		10: 10737418240,
		15: (15 * 1024 * 1024 * 1024),
	}
	for key, value := range examples {
		actual := util.BytesFromGb(key)
		if value != actual {
			t.Errorf("Expected '%d' to be '%d', but got '%d'", key, value, actual)
		}
	}
}
