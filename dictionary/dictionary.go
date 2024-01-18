package dictionary

import (
	"encoding/json"
	"errors"
	"os"
)

// Entry represents a dictionary entry with a word and its definition.
type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

// String returns the formatted representation of an Entry.
func (e Entry) String() string {
	return e.Word + ": " + e.Definition
}

// Dictionary represents the main structure of the dictionary,
// containing the filepath of the JSON file where entries are stored.
type Dictionary struct {
	filepath string
}

// New creates a new Dictionary instance with a given filepath.
func New(filepath string) *Dictionary {
	return &Dictionary{filepath: filepath}
}

// Add inserts a new word and its definition into the dictionary.
// Returns an error if the word already exists.
func (d *Dictionary) Add(word string, definition string) error {
	entries, err := d.readEntries()
	if err != nil {
		return err
	}

	if _, exists := entries[word]; exists {
		return errors.New("Word already exists")
	}

	entries[word] = Entry{Word: word, Definition: definition}
	return d.writeEntries(entries)
}

// Get retrieves an entry by word.
// Returns an error if the word does not exist in the dictionary.
func (d *Dictionary) Get(word string) (Entry, error) {
	entries, err := d.readEntries()
	if err != nil {
		return Entry{}, err
	}

	entry, exists := entries[word]
	if !exists {
		return Entry{}, errors.New("Word does not exist")
	}

	return entry, nil
}

// Remove deletes an entry by word.
// Returns an error if the word does not exist in the dictionary.
func (d *Dictionary) Remove(word string) error {
	entries, err := d.readEntries()
	if err != nil {
		return err
	}

	if _, exists := entries[word]; !exists {
		return errors.New("Word does not exist")
	}

	delete(entries, word)
	return d.writeEntries(entries)
}

// List returns a list of all words in the dictionary.
func (d *Dictionary) List() ([]string, error) {
	entries, err := d.readEntries()
	if err != nil {
		return nil, err
	}

	var words []string
	for word := range entries {
		words = append(words, word)
	}
	return words, nil
}

// readEntries reads and unmarshals the entries from the JSON file.
func (d *Dictionary) readEntries() (map[string]Entry, error) {
	file, err := os.ReadFile(d.filepath)
	if err != nil {
		if os.IsNotExist(err) || len(file) == 0 {
			return make(map[string]Entry), nil
		}
		return nil, err
	}

	var entries map[string]Entry = make(map[string]Entry)
	if len(file) > 0 {
		err = json.Unmarshal(file, &entries)
		if err != nil {
			return nil, err
		}
	}
	return entries, nil
}

// writeEntries marshals and writes the entries to the JSON file.
func (d *Dictionary) writeEntries(entries map[string]Entry) error {
	file, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	return os.WriteFile(d.filepath, file, 0666)
}
