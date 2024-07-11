package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Performs Anki search and returns the first note ID found
func FindNoteID(key string, deckName string) (int, error) {
	// Perform query
	params := map[string]interface{}{
		"query": "added:1 Sentence: Key:" + key + " deck:" + deckName, // Cards added today with an empty sentence field
	}
	fmt.Printf("Query: %v", params)
	res, err := InvokeAnkiRequest("findNotes", params)
	if err != nil {
		return -1, fmt.Errorf("Failed to search for note: %w", err)
	}

	// Decode response
	var result []int
	err = json.Unmarshal(res, &result)
	if err != nil {
		return -1, fmt.Errorf("Failed to unmarshal json: %w", err)
	}

	// Return first result upon success
	if len(result) == 0 {
		return -1, fmt.Errorf("No notes found matching search query %v", params)
	}
	id := result[0]
	return id, nil
}

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
		return json.RawMessage{}, fmt.Errorf("Could not marshal json: %w", err)
	}

	// Send request
	res, err := http.Post("http://127.0.0.1:8765", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return json.RawMessage{}, fmt.Errorf("Failed to send request: %w", err)
	}

	// Read response body
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return json.RawMessage{}, fmt.Errorf("Failed to read body of response: %w", err)
	}

	// Decode response
	var data *AnkiResponse
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		return json.RawMessage{}, fmt.Errorf("Failed to unmarshal json: %w", err)
	}

	// Return result upon success
	if data.Error != nil {
		return json.RawMessage{}, fmt.Errorf("Anki-Connect error: " + *data.Error)
	}
	return data.Result, nil
}