package models

// FormatType is the Event Source format type
// This should parse from the first line of source text
type FormatType int

const (
	// Traditional file format
	// Store all raw text on sqlite
	FormatMarkdown FormatType = iota
	FormatHTML
	FormatTXT

	// TODO:
	// experimental format link
	// fetch data by network
	// FormatLink

	// other format can be extend later
)
