# anki-immersion-reader

## Usage

The binary can be found [here](https://github.com/GameFuzzy/anki-immersion-reader/releases).

```
./anki-immersion-reader filepath [deck]
```

with `filepath` being the path to the CSV file exported by AnkiDojo and `deck` being the name of your mining deck in Anki.

If `deck` is omitted then deck:current is used.

## Build from source

### Build prerequisites
Go 1.22.0

### Instructions
```go build```
