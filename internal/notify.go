package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

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

var CellHeaderTemplate CellItem = CellItem{
	Wrap:   true,
	Type:   "TextBlock",
	Weight: "Bolder",
	Size:   "ExtraLarge",
	Style:  "heading",
}

var CellItemTemplate CellItem = CellItem{
	Type:   "TextBlock",
	Weight: "default",
	Size:   "small",
	Style:  "default",
	Wrap:   true,
}

func Notify(rows []TableRow) {

	jsonTemplateFile := "teams_file"

	teamsUrl := GetConfigValue("teams_url")
	jsonBody, err := openJsonFile(jsonTemplateFile)
	if err != nil {
		log.Fatalf("Failed to open json file... %v", err)
	}

	postBody, err := createTeamsMessage(jsonBody, rows)
	if err != nil {
		log.Fatalf("failed to create POST body... %s", err)
	}
	fmt.Printf("%s \n", postBody)
	sendTeamsMessage(teamsUrl, postBody)
}

func sendTeamsMessage(url string, postBody []byte) error {

	respBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", respBody)

	if err != nil {
		log.Fatalf("An error occurred sending to teams... %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An error occurred..s %v", err)
	}
	sb := string(body)

	// Just print for now
	fmt.Println(sb)
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
func createTeamsMessage(jsonObj map[string]any, rows []TableRow) (jsonMarshal []byte, err error) {

	// Get reference to rows

	attachmentsSlice, _ := jsonObj["attachments"].([]any)
	attachmentsObj, _ := attachmentsSlice[0].(map[string]any)
	content, _ := attachmentsObj["content"].(map[string]any)
	body, _ := content["body"].([]any)
	table, _ := body[0].(map[string]any)
	tableRows, _ := table["rows"].([]TableRow)
	for _, rowObj := range rows {
		fmt.Println(rowObj)
		tableRows = append(tableRows, rowObj)
	}
	table["rows"] = tableRows
	jsonMarshal, err = json.MarshalIndent(jsonObj, "", "  ")
	return jsonMarshal, err

}

// Takes the headers and list of lists with row values
// return final tablerows to be inserted to final json payload
func NotifyFormat(headers []string, items [][]string) (allRows []TableRow, err error) {

	// Create headers and add to allRows
	// Only multiple cells required
	headerRow := TableRow{
		Type:  "TableRow",
		Cells: []Cell{},
	}

	for _, v := range headers {
		cell := Cell{
			Type:  "TableCell",
			Items: []CellItem{},
		}
		item := CellHeaderTemplate
		item.Text = v
		cell.Items = append(cell.Items, item)
		headerRow.Cells = append(headerRow.Cells, cell)
	}

	allRows = append(allRows, headerRow)
	// Iterate through each array of items and add as a row, similar to headerRow
	for _, v := range items {
		dataRow := TableRow{
			Type:  "TableRow",
			Cells: []Cell{},
		}
		for _, cellItem := range v {
			cell := Cell{
				Type:  "TableCell",
				Items: []CellItem{},
			}
			item := CellItemTemplate
			item.Text = cellItem
			cell.Items = append(cell.Items, item)
			dataRow.Cells = append(dataRow.Cells, cell)
		}
		allRows = append(allRows, dataRow)
	}

	// Create data rows and add to allRows
	fmt.Println(allRows)
	return allRows, err
}
