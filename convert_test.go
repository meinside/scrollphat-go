package scroll

import (
	"testing"

	"bytes"
)

func TestArraysToBytes(t *testing.T) {
	// wrongly sized arrays
	if ArraysToBytes([][]uint8{{0, 0}, {1, 1}}) != nil {
		t.Errorf("Passed array should be sized as %dx%d.", NumRows, NumCols)
	}

	// hand-written arrays
	converted := ArraysToBytes([][]uint8{
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
	})
	expected := []byte{31, 0, 31, 0, 31, 0, 31, 0, 31, 0, 31}
	if !bytes.Equal(converted, expected) {
		t.Error("Arrays were not converted correctly.")
	}
}
