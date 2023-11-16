// Package dictionary
package dictionary

import (
	"encoding/json" // Used for encoding and decoding JSON data
	"errors"        // Library package for handling errors
	"os"            // For file system operations (read, write, etc.)
)

// Entry defines a dictionary entry with a string definition.
type Entry struct {
	Definition string `json:"definition"`
}

// String returns the definition of the entry as a string.
func (e Entry) String() string {
	return e.Definition
}

type Dictionary struct {
	filepath string // filepath is the path to the file where the dictionary is stored
}

// New init and returns a reference to a new Dictionary with a given filepath.
func New(filepath string) *Dictionary {
	return &Dictionary{filepath: filepath}
}

// Add inserts a word and its definition into the dictionary file.
func (d *Dictionary) Add(word string, definition string) error {
	entries, err := d.readEntries() // Retrieves the current entries from file
	if err != nil {
		return err
	}

	if _, exists := entries[word]; exists {
		return errors.New("Word already exists")
	}

	entries[word] = Entry{Definition: definition}
	return d.writeEntries(entries)
}

// Get retrieves a word's definition from the dictionary file.
func (d *Dictionary) Get(word string) (Entry, error) {
	entries, err := d.readEntries() // Retrieves the current entries from the file
	if err != nil {
		return Entry{}, err
	}

	entry, exists := entries[word]
	if !exists {
		return Entry{}, errors.New("Word does not exist")
	}

	return entry, nil
}

// Remove deletes a word from the dictionary file.
func (d *Dictionary) Remove(word string) error {
	entries, err := d.readEntries()
	if err != nil {
		return err
	}

	if _, exists := entries[word]; !exists {
		return errors.New("Word does not exist")
	}

	delete(entries, word)
	return d.writeEntries(entries) // Write and updated entries back to the file
}

// List returns all words and their definitions from the dictionary file.
func (d *Dictionary) List() ([]string, error) {
	entries, err := d.readEntries()
	if err != nil {
		return nil, err
	}

	var words []string
	for word := range entries {
		words = append(words, word) // Appends each word to the words slice
	}
	return words, nil
}

// ReadEntries reads and decodes entries from dictionary file.
func (d *Dictionary) readEntries() (map[string]Entry, error) {
	file, err := os.ReadFile(d.filepath)
	if err != nil {
		if os.IsNotExist(err) || len(file) == 0 {
			return make(map[string]Entry), nil // Initialise new map if file does not exist or is empty
		}
		return nil, err
	}

	var entries map[string]Entry = make(map[string]Entry) // Init map to store entries
	if len(file) > 0 {
		err = json.Unmarshal(file, &entries)
		if err != nil {
			return nil, err
		}
	}
	return entries, nil
}

// writeEntries encodes and writes entries to dictionary file.
func (d *Dictionary) writeEntries(entries map[string]Entry) error {
	file, err := json.Marshal(entries) // Encodes entries into JSON format
	if err != nil {
		return err
	}
	return os.WriteFile(d.filepath, file, 0666) // Writes encoded entries to file and returns any errors
}
