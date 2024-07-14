package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.design/x/clipboard"

	"encoding/csv"
)

var sentenceField string
var sentenceReadingField string

var filePath string
var deckName string

const (
	AnkiDojoWordColumn     = 1
	AnkiDojoSentenceColumn = 3
)

func init() {

	flag.Usage = func() {
		w := flag.CommandLine.Output()

		fmt.Fprintf(w, "Usage: %s [options] file_path deck\n", os.Args[0])

		flag.PrintDefaults()
	}

	const (
		defaultSentenceField        = "Sentence"
		defaultSentenceReadingField = "SentenceReading"
	)

	flag.StringVar(&sentenceField, "sentence", defaultSentenceField, "Name of your note type's sentence field")
	flag.StringVar(&sentenceField, "s", defaultSentenceField, "shorthand for -sentence")

	flag.StringVar(&sentenceReadingField, "sentence_reading", defaultSentenceReadingField, "Name of your note type's sentence reading field")
	flag.StringVar(&sentenceReadingField, "r", defaultSentenceReadingField, "shorthand for -sentence_reading")

	flag.Parse()

	// Retrieve arguments
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	filePath = flag.Arg(0)
	deckName = flag.Arg(1)
}

func main() {

	// Create map from words to sentences
	wordSentenceMap, err := createWordSentenceMapFromAnkiDojoExport()

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

	fmt.Println("Word list has been copied to the clipboard. Paste it into Yomitan's word generator and send the words to Anki.")
	fmt.Println("Press enter to continue once you've done the above...")
	// Wait for input
	fmt.Scanln()

	for word, sentence := range wordSentenceMap {
		id, err := FindNoteID(word, deckName)
		if err != nil {
			log.Printf("Failed to find note:\n%v\n", err)
			continue
		}

		if sentenceReadingField != "" {
			err = UpdateNoteFields(id, map[string]interface{}{
				sentenceField:        sentence,
				sentenceReadingField: "",
			})
		} else {
			err = UpdateNoteFields(id, map[string]interface{}{
				sentenceField: sentence,
			})
		}
		if err != nil {
			log.Printf("Failed to update note:\n%v\n", err)
		}
	}
}

// Returns map of words to sentences given the path of an AnkiDojo CSV file
func createWordSentenceMapFromAnkiDojoExport() (map[string]string, error) {
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
		m[row[AnkiDojoWordColumn]] = row[AnkiDojoSentenceColumn]
	}

	return m, nil
}
