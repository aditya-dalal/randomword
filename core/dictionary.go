package core

import (
	"os"
	"bufio"
)

type Dictionary interface {
	GetWords() []string
	LoadWords() error
}

type FileDictionary struct {
	File string
	Words map[string]bool
}

func (fd *FileDictionary) LoadWords() error {
	f, err := os.Open(fd.File)
	if err != nil {
		return err
	}
	fd.Words = make(map[string]bool)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fd.Words[scanner.Text()] = true
	}
	return nil
}

func (fd *FileDictionary) GetWords() []string {
	var words []string
	for word := range fd.Words {
		words = append(words, word)
	}
	return words
}