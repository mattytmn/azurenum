package internal

var (
	resultTable string
	tableEntry  []string
)

// Options for various output formats
func PrintResultAsTable() {}

func PrintResultAsJson() {}

func PrintResultAsCsv() {}

func PrintResultAsRaw() {}

// Takes in a slice and generates a table
// Use a 2D slice or struct to get header values and populate correct column
func createTableEntry(tableValues []*string) {}
