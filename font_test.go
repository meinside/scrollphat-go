package scroll

import (
	"testing"

	"bytes"
)

func TestTrim(t *testing.T) {
	var trimmed, expected []byte

	// all blank => two spaces
	trimmed = trim([]byte{0, 0, 0, 0, 0, 0})
	expected = []byte{0, 0}
	if !bytes.Equal(trimmed, expected) {
		t.Error("All-zero bytes should be converted to two zeros.")
	}

	// trim spaces on left side
	trimmed = trim([]byte{0, 0, 1, 1, 0, 0, 1, 1})
	expected = []byte{1, 1, 0, 0, 1, 1}
	if !bytes.Equal(trimmed, expected) {
		t.Error("Trimming on left side failed.")
	}

	// trim spaces on right side
	trimmed = trim([]byte{1, 1, 0, 0, 1, 1, 0, 0})
	expected = []byte{1, 1, 0, 0, 1, 1}
	if !bytes.Equal(trimmed, expected) {
		t.Error("Trimming on right side failed.")
	}

	// trim spaces on both sides
	trimmed = trim([]byte{0, 0, 1, 1, 0, 0, 1, 1, 0, 0})
	expected = []byte{1, 1, 0, 0, 1, 1}
	if !bytes.Equal(trimmed, expected) {
		t.Error("Trimming on both sides failed.")
	}
}

func TestBytesForCharacter(t *testing.T) {
	var converted, expected []byte

	// unconvertible character => spaces
	converted = BytesForCharacter('헐')
	expected = []byte{0, 0}
	if !bytes.Equal(converted, expected) {
		t.Error("Unconvertible character should be converted to spaces.")
	}

	// convertible character
	converted = BytesForCharacter('a')
	expected = []byte{8, 20, 28}
	if !bytes.Equal(converted, expected) {
		t.Error("Failed converting a character.")
	}
}

func TestBytesForString(t *testing.T) {
	var converted, expected []byte

	// blank string
	converted = BytesForString("    ")
	expected = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !bytes.Equal(converted, expected) {
		t.Error("Failed converting a blank string.")
	}

	// ordinary string
	converted = BytesForString("Go go go.")
	expected = []byte{14, 17, 29, 0, 8, 20, 8, 0, 0, 0, 0, 2, 21, 15, 0, 8, 20, 8, 0, 0, 0, 0, 2, 21, 15, 0, 8, 20, 8, 0, 16, 0}
	if !bytes.Equal(converted, expected) {
		t.Error("Failed converting a string.")
	}
}
