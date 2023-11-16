package dictionary

import "errors"

type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return e.Definition
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
	return &Dictionary{entries: make(map[string]Entry)}
}

func (d *Dictionary) Add(word string, definition string) error {
	if _, exists := d.entries[word]; exists {
		return errors.New("word already exists")
	}

	d.entries[word] = Entry{Definition: definition}
	return nil
}

func (d *Dictionary) Get(word string) (Entry, error) {
	entry, exists := d.entries[word]
	if !exists {
		return Entry{}, errors.New("word does not exist")
	}

	return entry, nil
}

func (d *Dictionary) Remove(word string) error {
	if _, exists := d.entries[word]; !exists {
		return errors.New("word does not exist")
	}

	delete(d.entries, word)
	return nil
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
	var words []string
	for word := range d.entries {
		words = append(words, word)
	}
	return words, d.entries
}
