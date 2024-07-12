package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.design/x/clipboard"

	"encoding/csv"
	"encoding/json"
)

func main() {
	// Retrieve arguments
	if len(os.Args) < 3 {
		panic(errors.New("Not enough arguments provided (expected ./anki-immersion-reader filepath deck [field]"))
	}

	filePath := os.Args[1]
	deckName := os.Args[2]

	fieldName := "Sentence"
	if len(os.Args) > 3 {
		fieldName = os.Args[3]
	}

	// Create map from words to sentences
	wordSentenceMap, err := createWordSentenceMapFromAnkiDojoExport(filePath)

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
		id, err := findNoteID(word, deckName, fieldName)
		if err != nil {
			log.Printf("Failed to find note:\n%v\n", err)
			continue
		}

		err = updateNoteSentence(id, sentence)
		if err != nil {
			log.Printf("Failed to update note:\n%v\n", err)
		}
	}
}

// Updates sentence field of the Anki note with the given ID
func updateNoteSentence(id int, sentence string) error {
	params := map[string]interface{}{
		"note": map[string]interface{}{
			"id": id,
			"fields": map[string]interface{}{
				"Sentence": sentence,
			},
		},
	}
	_, err := InvokeAnkiRequest("updateNoteFields", params)
	return err
}

// Performs Anki search and returns the first note ID found
func findNoteID(key, deckName, fieldName string) (int, error) {
	// Perform query
	query := fmt.Sprintf("%s:<b>%s</b> OR %s:%s OR (Word:%s Sentence:) is:new deck:%s", fieldName, key, fieldName, key, key, deckName)
	params := map[string]interface{}{
		"query": query,
	}
	res, err := InvokeAnkiRequest("findNotes", params)
	if err != nil {
		return -1, fmt.Errorf("Failed to search for note: %w\n", err)
	}

	// Decode response
	var result []int
	err = json.Unmarshal(res, &result)
	if err != nil {
		return -1, fmt.Errorf("Failed to unmarshal json: %w\n", err)
	}

	// Return first result upon success
	if len(result) == 0 {
		return -1, fmt.Errorf("No notes found matching search query \"%s\"\n", query)
	}
	id := result[0]
	return id, nil
}

func createWordSentenceMapFromAnkiDojoExport(filePath string) (map[string]string, error) {
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
