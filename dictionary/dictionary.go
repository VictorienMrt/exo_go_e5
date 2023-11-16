package dictionary

import (
	"errors"
)

// Entry structure
type Entry struct {
	Definition string
}

// String representation of an Entry
func (e Entry) String() string {
	return e.Definition
}

// Dictionary structure
type Dictionary struct {
	entries map[string]Entry
}

// New creates a new Dictionary
func New() *Dictionary {
	return &Dictionary{entries: make(map[string]Entry)}
}

// Add a word and its definition to the dictionary
func (d *Dictionary) Add(word string, definition string) {
	d.entries[word] = Entry{Definition: definition}
}

// Get the definition of a word
func (d *Dictionary) Get(word string) (Entry, error) {
	entry, exists := d.entries[word]
	if !exists {
		return Entry{}, errors.New("word not found")
	}
	return entry, nil
}

// Remove a word from the dictionary
func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

// List all words in the dictionary
func (d *Dictionary) List() ([]string, map[string]Entry) {
	var words []string
	for word := range d.entries {
		words = append(words, word)
	}
	return words, d.entries
}
	