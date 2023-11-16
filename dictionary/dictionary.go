// Package dictionary
package dictionary

import "errors"

// Entry defines a dictionary entry with a string definition.
type Entry struct {
	Definition string
}

// String returns the definition of the entry.
func (e Entry) String() string {
	return e.Definition
}

// Dictionary represents a collection of dictionary entries.
type Dictionary struct {
	entries map[string]Entry
}

// New initializes and returns a new Dictionary.
func New() *Dictionary {
	return &Dictionary{entries: make(map[string]Entry)}
}

// Add inserts a word and its definition into the dictionary.
func (d *Dictionary) Add(word string, definition string) error {
	if _, exists := d.entries[word]; exists {
		return errors.New("word already exists")
	}
	d.entries[word] = Entry{Definition: definition}
	return nil
}

// Get retrieves a word's definition from the dictionary.
func (d *Dictionary) Get(word string) (Entry, error) {
	entry, exists := d.entries[word]
	if !exists {
		return Entry{}, errors.New("word does not exist")
	}
	return entry, nil
}

// Remove deletes a word from the dictionary.
func (d *Dictionary) Remove(word string) error {
	if _, exists := d.entries[word]; !exists {
		return errors.New("word does not exist")
	}
	delete(d.entries, word)
	return nil
}

// List returns all words and their definitions from the dictionary.
func (d *Dictionary) List() ([]string, map[string]Entry) {
	var words []string
	for word := range d.entries {
		words = append(words, word)
	}
	return words, d.entries
}
