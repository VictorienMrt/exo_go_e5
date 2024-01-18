package dictionary

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Function Test Post (Add)
func TestAdd(t *testing.T) {
	// Create a new dictionary
	d := New("test.json")
	defer os.Remove("test.json") // Cleanup after the test

	// Test adding a new word
	err := d.Add("hello", "hello there")
	assert.Nil(t, err, "Add should not return an error for a new word")

	// Test adding a duplicate word
	err = d.Add("hello", "hello there")
	assert.NotNil(t, err, "Add should return an error for a duplicate word")
}

func TestGet(t *testing.T) {
	d := New("test.json")
	defer os.Remove("test.json")

	d.Add("hello", "hello there")

	// Test retrieving an existing word
	entry, err := d.Get("hello")
	assert.Nil(t, err, "Get should not return an error for an existing word")
	// Check the definition
	assert.Equal(t, "hello there", entry.Definition, "Get should return the correct definition")

	// Test retrieving a non-existing word
	_, err = d.Get("world")
	assert.NotNil(t, err, "Get should return an error for a non-existing word")
}

func TestRemove(t *testing.T) {
	d := New("test.json")
	defer os.Remove("test.json")

	d.Add("hello", "hello there")

	// Test removing an existing word
	err := d.Remove("hello")
	assert.Nil(t, err, "Remove should not return an error for an existing word")

	// Test removing a non-existing word
	err = d.Remove("world")
	assert.NotNil(t, err, "Remove should return an error for a non-existing word")
}

func TestList(t *testing.T) {
	d := New("test.json")
	defer os.Remove("test.json")

	d.Add("hello", "hello there")

	// Test listing entries
	entries, err := d.List() // Get all entries
	assert.Nil(t, err, "List should not return an error")
	// Check length of the list - minimum 1
	assert.Greater(t, len(entries), 0, "List should return at least one entry")
}
