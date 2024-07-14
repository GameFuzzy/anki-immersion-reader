package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AnkiRequest struct {
	Action  string                 `json:"action"`
	Params  map[string]interface{} `json:"params"`
	Version int                    `json:"version"`
}

type AnkiResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *string         `json:"error"`
}

func CreateAnkiRequest(action string, params map[string]interface{}) AnkiRequest {
	return AnkiRequest{
		Action:  action,
		Params:  params,
		Version: 6,
	}
}

func InvokeAnkiRequest(action string, params map[string]interface{}) (json.RawMessage, error) {
	// Create request
	jsonData, err := json.Marshal(CreateAnkiRequest(action, params))
	if err != nil {
		return json.RawMessage{}, fmt.Errorf("Could not marshal json: %w\n", err)
	}

	// Send request
	res, err := http.Post("http://127.0.0.1:8765", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return json.RawMessage{}, fmt.Errorf("Failed to send request: %w\n", err)
	}

	// Read response body
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return json.RawMessage{}, fmt.Errorf("Failed to read body of response: %w\n", err)
	}

	// Decode response
	var data *AnkiResponse
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		return json.RawMessage{}, fmt.Errorf("Failed to unmarshal json: %w\n", err)
	}

	// Return result upon success
	if data.Error != nil {
		return json.RawMessage{}, fmt.Errorf("Anki-Connect error: %s\n", *data.Error)
	}
	return data.Result, nil
}

// Updates fields of the Anki note with the given ID
func UpdateNoteFields(id int, fields map[string]interface{}) error {
	params := map[string]interface{}{
		"note": map[string]interface{}{
			"id":     id,
			"fields": fields,
		},
	}
	_, err := InvokeAnkiRequest("updateNoteFields", params)
	return err
}

// Performs Anki search and returns the first note ID found
func FindNoteID(key, deckName string) (int, error) {
	// Perform query
	query := fmt.Sprintf("%s:<b>%s</b> OR %s:%s OR (Word:%s %s:) is:new deck:%s", sentenceField, key, sentenceField, key, key, sentenceField, deckName)
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
