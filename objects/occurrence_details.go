package objects

// Code generated from OCSF schema; DO NOT EDIT.

// OccurrenceDetails Details about where in the target entity, specified information was discovered.
// Only the attributes, relevant to the target entity type should be populated.
type OccurrenceDetails struct {
	// The cell name/reference in a spreadsheet. e.g A2
	CellName string `json:"cell_name,omitempty"`

	// The column name in a spreadsheet, where the information was discovered.
	ColumnName string `json:"column_name,omitempty"`

	// The column number in a spreadsheet or a plain text document, where the
	// information was discovered.
	ColumnNumber int64 `json:"column_number,omitempty"`

	// The line number of the last line of the file, where the information was
	// discovered.
	EndLine int64 `json:"end_line,omitempty"`

	// The JSON path of the attribute in a json record, where the information was
	// discovered
	JsonPath string `json:"json_path,omitempty"`

	// The page number in a document, where the information was discovered.
	PageNumber int64 `json:"page_number,omitempty"`

	// The index of the record in the array of records, where the information was
	// discovered. e.g. the index of a record in an array of JSON records in a
	// file.
	RecordIndexInArray int64 `json:"record_index_in_array,omitempty"`

	// The row number in a spreadsheet, where the information was discovered.
	RowNumber int64 `json:"row_number,omitempty"`

	// The line number of the first line of the file, where the information was
	// discovered.
	StartLine int64 `json:"start_line,omitempty"`
}
