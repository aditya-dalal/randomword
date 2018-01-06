package main

import (
	"testing"
	"randomword/core"
	"errors"
)

func TestGetRandomWords(t *testing.T) {
	tests := []struct{
		num int
		result int
		err error
	}{
		{2, 2, nil},
		{-1, 0, nil},
		{6, 0, errors.New("Random number count is greater than available words")},
	}
	for _, test := range tests {
		var dictionary core.Dictionary
		dictionary = &TestDictionary{}
		words, err := GetRandomWords(dictionary, test.num)
		if len(words) != test.result {
			t.Errorf("Expected %d words, actual %d words", test.num, len(words))
		}
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected '%v', actual '%v'", test.err, err)
		}
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