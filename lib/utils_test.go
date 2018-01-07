package lib

import (
	"testing"
	"time"
	"math/rand"
	"errors"
)

func TestGetRandomInt(t *testing.T) {
	tests := []struct{
		min int
		max int
		err error
	} {
		{0, 1, nil},
		{-5, 5, nil},
		{0, 0, errors.New("min value should be greater than max value")},
	}

	rand.Seed(time.Now().Unix())
	for _, test := range tests {
		num, err := GetRandomInt(test.min, test.max)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Expected: %v, actual: %v", test.err, err)
		}
		if err == nil && (num < test.min || num >= test.max) {
			t.Errorf("Expected: %d in between %d and %d", num, test.min, test.max)
		}
	}
}