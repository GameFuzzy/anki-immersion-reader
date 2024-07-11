package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.design/x/clipboard"

	"encoding/csv"
)

func main() {
	// Retrieve arguments
	if len(os.Args) < 2 {
		panic(errors.New("Not enough arguments provided (expected path to AnkiDojo export file"))
	}

	deckName := "current"
	if len(os.Args) > 2 {
		deckName = os.Args[2]
	}

	// Create map from words to sentences
	wordSentenceMap, err := CreateWordSentenceMapFromAnkiDojoExport(os.Args[1])

	// Make array of just words
	keys := make([]string, 0, len(wordSentenceMap))
	for k := range wordSentenceMap {
		keys = append(keys, k)
	}

	// Copy list of words separated by newlines to clipboard
	err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(strings.Join(keys, "\n")))

	fmt.Println("Word list has been copied to the clipboard. Paste it into Yomitan's word generator and send the words to Anki")
	fmt.Println("Press enter to continue once you've done the above...")
	// Wait for input
	fmt.Scanln()

	for word, sentence := range wordSentenceMap {
		id, err := FindNoteID(word, deckName)
		if err != nil {
			log.Printf("Failed to find note: %s\n", err)
			continue
		}

		params := map[string]interface{}{
			"note": map[string]interface{}{
				"id": id,
				"fields": map[string]interface{}{
					"Sentence": sentence,
				},
			},
		}
		_, err = InvokeAnkiRequest("updateNoteFields", params)
		if err != nil {
			log.Printf("Failed to update note: %s\n", err)
		}
	}
}

func CreateWordSentenceMapFromAnkiDojoExport(filePath string) (map[string]string, error) {
	// Read file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	r := csv.NewReader(f)

	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	data = data[1:] // Discard first row

	// Create map
	m := make(map[string]string)
	for _, row := range data {
		m[row[0]] = row[3]
	}

	return m, nil
}
