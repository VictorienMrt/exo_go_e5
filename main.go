// main package definition
package main

// Importing packages
import (
	"bufio" // For reading input
	"exo_go_e5/dictionary"
	"fmt"     // For formatted I/O operations
	"os"      // For interfacing with the operating system
	"strings" // For string manipulation
)

// Function main
func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader for standard input
	d := dictionary.New("test.json")    // Init a new dictionary object

	// Infinite loop to handle user input
	for {
		fmt.Println("\nSelect an action [add, define, remove, list, exit]:")
		action, _ := reader.ReadString('\n') // Action choice from user
		action = strings.TrimSpace(action)

		// Switch statement to handle the selected action
		switch action {
		case "add":
			actionAdd(d, reader) // Calling the corresponding function
		case "define":
			actionDefine(d, reader)
		case "remove":
			actionRemove(d, reader)
		case "list":
			actionList(d)
		case "exit":
			fmt.Println("End of program")
			return
		default:
			fmt.Println("Not recognized.") // Print error
		}
	}
}

// Function to add a word to the dictionary
func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	fmt.Print("Enter a definition: ")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)

	err := d.Add(word, definition) // Add the word to the dictionary
	if err != nil {
		fmt.Println("Failed to add word:", err) // Print error
	} else {
		fmt.Println("Word added.") // Print success message
	}
}

// Function to define a word in the dictionary
func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	entry, err := d.Get(word)
	if err != nil {
		fmt.Println("Failed to find word:", err)
	} else {
		fmt.Println("Definition:", entry)
	}
}

// Function to remove a word from the dictionary
func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to remove: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	err := d.Remove(word) // Remove the word from the dictionary
	if err != nil {
		fmt.Println("Failed to remove word:", err)
	} else {
		fmt.Println("Word removed.")
	}
}

// Function to list all words and their definitions in the dictionary
func actionList(d *dictionary.Dictionary) {
	words, err := d.List() // Attempt to get the list of words from the dictionary
	if err != nil {
		fmt.Println("Failed to list words:", err)
		return
	}

	for _, word := range words {
		entry, err := d.Get(word) // For each word, attempt to get the corresponding entry
		if err != nil {
			fmt.Println("Failed to get definition for word:", word, "with error:", err)
			continue
		}
		fmt.Println("Word :", word, "--", "Description :", entry) // Print the word and its definition
	}
}
