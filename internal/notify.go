package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type MessageBody struct {
	Title string `json:"message_title,omitempty"`
	Body  string `json:"message_body,omitempty"`
}

type TableRow struct {
	Type  string `json:"type" default:"TableRow"`
	Cells []Cell `json:"cells"`
}

type Cell struct {
	Type  string     `json:"type" default:"TableCell"`
	Items []CellItem `json:"items"`
}

type HeaderCellItem struct {
	Type   string `json:"type" default:"TextBlock"`
	Text   string `json:"text"`
	Weight string `json:"weight" default:"Bolder"`
	Size   string `json:"size" default:"ExtraLarge"`
	Style  string `json:"style" default:"heading"`
	Wrap   bool   `json:"wrap"`
}

type CellItem struct {
	Type   string `json:"type"`
	Text   string `json:"text"`
	Size   string `json:"size"`
	Style  string `json:"style"`
	Weight string `json:"weight"`
	Wrap   bool   `json:"wrap"`
}

func SendTeamsMessage() error {
	jsonTemplateFile := "template.json"

	teamsUrl := GetConfigValue("teams_url")
	jsonBody, err := openJsonFile(jsonTemplateFile)
	if err != nil {
		log.Fatalf("Failed to open json file... %v", err)
	}
	postBody, err := createTeamsMessage(jsonBody)
	if err != nil {
		log.Fatalf("failed to create POST body... %s", err)
	}

	respBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(teamsUrl, "application/json", respBody)

	if err != nil {
		log.Fatalf("An error occurred sending to teams... %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An error occurred..s %v", err)
	}
	sb := string(body)
	log.Println(sb)
	return err
}

// Given a json template, return unmarshalled
func openJsonFile(filename string) (map[string]any, error) {

	jsonFilePath := GetConfigValue(filename)

	jsonString, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("An error occurred while opening JSON file... %s", err)
	}
	var jsonUnmarshalled map[string]any

	err = json.Unmarshal(jsonString, &jsonUnmarshalled)
	if err != nil {
		log.Fatalf("failed to unmarshal json file... %s", err)
	}

	return jsonUnmarshalled, err
}

// Takes in JSON template and adds custom headers and rows to table
func createTeamsMessage(jsonUnmarshahlled map[string]any) (jsonMarshalled []byte, err error) {

	return jsonMarshalled, err

}
