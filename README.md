# anki-immersion-reader

[![Create release](https://github.com/GameFuzzy/anki-immersion-reader/actions/workflows/go.yml/badge.svg)](https://github.com/GameFuzzy/anki-immersion-reader/actions/workflows/go.yml)

## Description

### What the script does

1. Copies a list of all the words in the provided CSV file exported by Immersion Reader and prompts you to paste them into Yomitan's word generator and hit generate.
2. Waits for the user to hit enter to continue.
3. Once the user has done so it adds the sentences extracted from the CSV file to the newly created Anki cards.
4. (Optional) It will then by default clear the sentence reading field to allow for easier [backfilling of furigana](https://arbyste.github.io/jp-mining-note-prerelease/faq/#how-do-i-bulk-generate-furigana-and-pitch-accents).

## Installation

Prebuilt binaries for most platforms can be found [here](https://github.com/GameFuzzy/anki-immersion-reader/releases/latest).

You can also build it from source as described [further down on this page](#build-from-source).

### (Optional) Verify binary checksum

Requires [GitHub CLI](https://cli.github.com)

```sh
gh attestation verify anki-immersion-reader_<OS>_<ARCH>[.exe] -R GameFuzzy/anki-immersion-reader
```

## Usage

```text
./anki-immersion-reader_<OS>_<ARCH>[.exe] [options] filepath deck
  -r string
        shorthand for -sentence_reading (default "SentenceReading")
  -s string
        shorthand for -sentence (default "Sentence")
  -sentence string
        Name of your note type's sentence field (default "Sentence")
  -sentence_reading string
        Name of your note type's sentence reading field (default "SentenceReading")
```

with `filepath` being the path to the CSV file exported by AnkiDojo,

`deck` being the name of your mining deck in Anki.

### Notes

- Setting `sentence_reading` to an empty string like so: `-sentence_reading ""` tells the script to not clear the sentence reading field.

## Build from source

Requires Go 1.22.0

### Instructions

```sh
go build
```
