package dictionary

import (
	"encoding/json"
	"errors"
	"os"
)

type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

func (e Entry) String() string {
	return e.Word + ": " + e.Definition
}

type Dictionary struct {
	filepath string
}

func New(filepath string) *Dictionary {
	return &Dictionary{filepath: filepath}
}

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

func (d *Dictionary) writeEntries(entries map[string]Entry) error {
	file, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	return os.WriteFile(d.filepath, file, 0666)
}
