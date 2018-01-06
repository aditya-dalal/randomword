package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"math/rand"
	"time"
	"errors"
	"randomword/core"
	"randomword/lib"
)

const FILE = "/tmp/words_alpha.txt"

func GetDictionary() (*core.FileDictionary, error) {
	dictionary := &core.FileDictionary{File: FILE}
	err := dictionary.LoadWords()
	if err != nil {
		return nil, err
	}
	return dictionary, nil
}

func GetRandomWords(dictionary core.Dictionary, num int) ([]string, error) {
	words := dictionary.GetWords()
	if num >= len(words) {
		return nil, errors.New("Random number count is greater than available words")
	}
	var result []string
	for i := 0; i < num; i++ {
		index := lib.GetRandomInt(i, len(words))
		words[i], words[index] = words[index], words[i]
		result = append(result, words[i])
	}
	return result, nil
}

func PrintWords(words []string) {
	for _, word := range words {
		fmt.Println(word)
	}
}

func ParseParameter() (int, error) {
	var num int
	var err error
	if len(os.Args) < 2 {
		num = 1
	} else {
		num, err = strconv.Atoi(os.Args[1])
		if err != nil || num < 1 {
			return 0, errors.New("Please provide positive integer as parameter")
		}
	}
	return num, nil
}

func main() {
	rand.Seed(time.Now().Unix())
	num, err := ParseParameter()
	if err != nil {
		log.Fatal(err)
	}
	var dictionary core.Dictionary
	dictionary, err = GetDictionary()
	if err != nil {
		log.Fatal(err)
	}
	words, err := GetRandomWords(dictionary, num)
	if err != nil {
		log.Fatal(err)
	}
	PrintWords(words)
}