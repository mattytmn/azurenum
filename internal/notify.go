package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var WEBHOOK string = "https://prod-02.australiasoutheast.logic.azure.com:443/workflows/e60c2e89b72b458f9a2ff0b37affec4a/triggers/manual/paths/invoke?api-version=2016-06-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=y8XAivf3HU1EXm7dQK0WNss78eM8wBGGIdOZ9HHaTrM"

type MessageBody struct {
	Title string `json:"message_title,omitempty"`
	Body  string `json:"message_body,omitempty"`
}

func (m *MessageBody) SendTeamsMessage() {
	jsonTemplateFile := "template.json"

	jsonBody := openJsonFile(jsonTemplateFile)
	customizeJson(jsonBody)
}

func openJsonFile(filename string) (jsonString []byte) {

	jsonString, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("An error occurred while opening JSON file... %s", err)
	}

	return jsonString
}

func customizeJson(jsonData []byte) error {

	var jsonObject map[string]any

	err := json.Unmarshal(jsonData, &jsonObject)
	if err != nil {
		log.Fatalf("An error occurred marshalling JSON.. %s", err)
	}

	for k, v := range jsonObject {
		fmt.Printf("k: %s, v: %s", k, v)
	}
	return nil

}
