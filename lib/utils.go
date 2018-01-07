package lib

import (
	"math/rand"
	"errors"
)

func GetRandomInt(min int, max int) (int, error) {
	if min >= max {
		return -1, errors.New("min value should be greater than max value")
	}
	return rand.Intn(max-min) + min, nil
}
