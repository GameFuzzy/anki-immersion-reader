# anki-immersion-reader

![Latest build](https://github.com/GameFuzzy/anki-immersion-reader/actions/workflows/go.yml/badge.svg)

## Description

### What the script does:
1. Copies a list of all the words in the provided CSV file exported by Immersion Reader and prompts you to paste them into Yomitan's word generator and hit generate.
2. Waits for the user to hit enter to continue.
3. Once the user has done so it adds the sentences extracted from the CSV file to the newly created Anki cards.

## Usage

Prebuilt binaries for most platforms can be found [here](https://github.com/GameFuzzy/anki-immersion-reader/releases/latest). You can also build it from source as described [further down on this page](#build-from-source).

```sh
./anki-immersion-reader_<OS>_<ARCH>[.exe] filepath deck [field]
```

with `filepath` being the path to the CSV file exported by AnkiDojo,

`deck` being the name of your mining deck in Anki,

and `field` being the name of your note type's sentence field.

### Notes:

- `field` will default to "Sentence" if omitted.

## Build from source

### Build prerequisites:

Go 1.22.0

### Instructions:

```go build```
