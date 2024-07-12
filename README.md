# anki-immersion-reader

## Usage

The binary can be found [here](https://github.com/GameFuzzy/anki-immersion-reader/releases).

```sh
./anki-immersion-reader filepath deck [field]
```

with `filepath` being the path to the CSV file exported by AnkiDojo,

`deck` being the name of your mining deck in Anki,

and `field` being the name of your note type's sentence field.

### Notes

- `field` will default to "Sentence" if omitted

## Build from source

### Build prerequisites

Go 1.22.0

### Instructions

```go build```
