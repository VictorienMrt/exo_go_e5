package main

import (
	"bufio"
	"exo_go_e5/dictionary" // assurez-vous que ce chemin correspond au nom du module dans go.mod
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	d := dictionary.New()

	for {
		fmt.Println("\nSelect an action [add, define, remove, list, exit]:")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		switch action {
		case "add":
			actionAdd(d, reader)
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
			fmt.Println("Not recognized.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	fmt.Print("Enter a definition: ")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition)

	err := d.Add(word, definition)
	if err != nil {
		fmt.Println("Failed to add word:", err)
	} else {
		fmt.Println("Word added.")
	}
}

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

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {
	fmt.Print("Enter a word to remove: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	err := d.Remove(word)
	if err != nil {
		fmt.Println("Failed to remove word:", err)
	} else {
		fmt.Println("Word removed.")
	}
}

func actionList(d *dictionary.Dictionary) {
	words, entries := d.List()
	for _, word := range words {
		entry := entries[word]
		fmt.Println(word, ":", entry)
	}
}
