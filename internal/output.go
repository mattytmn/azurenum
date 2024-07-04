package internal

import (
	"os"

	"github.com/aquasecurity/table"
)

var (
	resultTable string
	tableHead   []string
	tableBody   [][]string
)

// Options for various output formats

type TableClient struct {
	Header []string
	Body   [][]string
}

// Create new struct for table
func (t *TableClient) PrintResultAsTable(tblEntry TableClient) {
	// headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	tbl := table.New(os.Stdout)
	tbl.SetHeaders(tblEntry.Header...)
	for _, r := range tblEntry.Body {
		tbl.AddRow(r...)
	}

	tbl.Render()

}

func PrintResultAsJson() {}

func PrintResultAsCsv() {}

func PrintResultAsRaw() {}

// Takes in a slice and generates a table
// Use a 2D slice or struct to get header values and populate correct column
func createTableEntry(tableValues []*string) {}
