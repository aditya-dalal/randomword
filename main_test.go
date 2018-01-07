package main

import (
	"testing"
	"randomword/core"
	"errors"
)

func TestGetRandomWords(t *testing.T) {
	tests := []struct{
		dictionary core.Dictionary
		num int
		result int
		err error
	}{
		{&TestDictionary{}, 2, 2, nil},
		{&TestDictionary{}, -1, 0, nil},
		{&TestDictionary{}, 6, 0, errors.New("Random number count is greater than available words")},
		{&EmptyDictionary{}, 1, 0, errors.New("No words in dictionary")},
	}
	for _, test := range tests {
		words, err := GetRandomWords(test.dictionary, test.num)
		if len(words) != test.result {
			t.Errorf("Expected %d words, actual %d words", test.num, len(words))
		}
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected '%v', actual '%v'", test.err, err)
		}
	}
}

func TestGetDictionary(t *testing.T) {
	tests := []struct{
		file string
		err error
	} {
		{"words.txt", nil},
		{"abc", errors.New("open abc: no such file or directory")},
	}

	for _, test := range tests {
		dictionary, err := GetDictionary(test.file)
		if err == nil && dictionary == nil {
			t.Error("Ecpected dictionary to be not nil")
		}
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected: %v, actual: %v", test.err, err)
		}
	}
}

func TestParseParameter(t *testing.T) {
	tests := []struct{
		param string
		result int
		err error
	} {
		{"2", 2, nil},
		{"2.0", 0, errors.New("Please provide positive integer as parameter")},
		{"-2", 0, errors.New("Please provide positive integer as parameter")},
		{"a", 0, errors.New("Please provide positive integer as parameter")},
	}

	for _, test := range tests {
		val, err := ParseParameter([]string{"abc", test.param})
		if err == nil && val != test.result {
			t.Errorf("Expected: %d, actual: %d", test.result, val)
		}
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected: %v, actual: %v", test.err, err)
		}
	}

	val, _ := ParseParameter([]string{"abc"})
	if val != 1 {
		t.Errorf("Expected: %d, actual: %d", 1, val)
	}
}

type TestDictionary struct {
}

func (td *TestDictionary) GetWords() []string {
	return []string{"redefinitions", "uncredibly", "carven", "confinable", "handycuff"}
}

func (td *TestDictionary) LoadWords() error {
	return nil
}

type EmptyDictionary struct {
}

func (ed *EmptyDictionary) GetWords() []string {
	return nil
}

func (ed *EmptyDictionary) LoadWords() error {
	return nil
}