# anki-immersion-reader

## Usage

The binary can be found [here](https://github.com/GameFuzzy/anki-immersion-reader/releases).

```sh
./anki-immersion-reader filepath [deck] [field]
```

with `filepath` being the path to the CSV file exported by AnkiDojo,
`deck` being the name of your mining deck in Anki,
and

### Notes

If `deck` is omitted then "deck:current" will be used.
If `field` is omitted then the sentence field will be assumed to be "Sentence".

## Build from source

### Build prerequisites

Go 1.22.0

### Instructions

```go build```
