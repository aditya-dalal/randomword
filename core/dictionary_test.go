package core

import (
	"testing"
	"errors"
)

func TestFileDictionary_LoadWords(t *testing.T) {
	tests := []struct{
		file string
		err error
	} {
		{"../words.txt", nil},
		{"abc", errors.New("open abc: no such file or directory")},
	}

	for _, test := range tests {
		fd := &FileDictionary{File: test.file}
		err := fd.LoadWords()
		if err == nil {
			if fd.Words == nil {
				t.Errorf("Expected dictionary to be not nil, actual %v", fd.Words)
			} else if len(fd.Words) != 4 {
				t.Errorf("Expected 4 words map, actual %d words map", len(fd.Words))
			}
		} else if err.Error() != test.err.Error() {
			t.Errorf("Expected %v, actual %v", test.err.Error(), err.Error())
		}
	}
}

func TestFileDictionary_GetWords(t *testing.T) {
	tests := []struct{
		data     map[string]bool
		dataSize int
	} {
		{map[string]bool{"redefinitions":true, "uncredibly":true}, 2},
		{nil, 0},
	}
	for _, test := range tests {
		fd := &FileDictionary{Words: test.data}
		words := fd.GetWords()
		if len(words) != test.dataSize {
			t.Errorf("Expected: %d words, actual: %d words", test.dataSize, len(words))
		}
	}
}