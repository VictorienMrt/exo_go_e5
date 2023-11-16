package main

import (
	"estiam/dictionary"
	"fmt"
	"sort"
)

func main() {
	d := dictionary.New()

	// Add words and definitions
	d.Add("Go", "Test")
	d.Add("Map", "")
	d.Add("Function", "ok")

	// Display the definition of a word
	definition, err := d.Get("Go")
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("Go:", definition)
	}

	// Remove a word
	d.Remove("")

	// Display the sorted list of words and their definitions
	words, entries := d.List()
	sort.Strings(words)
	fmt.Println("\nDictionnaire:")
	for _, word := range words {
		fmt.Printf("%s: %s\n", word, entries[word].Definition)
	}
}
