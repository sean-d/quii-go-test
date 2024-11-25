package maps

import (
	"errors"
)

var ErrNotFound = errors.New("word not in dictionary")
var ErrWordExists = errors.New("word already in dictionary")

type Dictionary map[string]string

func (d Dictionary) Add(word, def string) error {
	_, ok := d[word]
	if !ok {
		d[word] = def
		return nil
	} else {
		return ErrWordExists
	}
}

func (d Dictionary) Update(word, newDef string) error {
	_, ok := d[word]

	if !ok {
		return ErrNotFound
	} else {
		d[word] = newDef
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, ok := d[word]

	if !ok {
		return ErrNotFound
	} else {
		delete(d, word)
	}

	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
