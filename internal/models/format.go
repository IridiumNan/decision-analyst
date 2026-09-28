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

	// Invalid format
	FormatInvalid

	// other format can be extend later
)

// FormatTodType convert the format (string on database) into [FormatType]
func FormatTodType(formatStr string) FormatType {
	switch formatStr {
	case "markdown":
		return FormatMarkdown
	case "html":
		return FormatHTML
	case "txt":
		return FormatTXT
	}

	return FormatInvalid
}

// FormatToText convert the [FormatType] into string like markdown html, etc...
func FormatToText(t FormatType) string {
	switch t {
	case FormatMarkdown:
		return "markdown"
	case FormatHTML:
		return "html"
	case FormatTXT:
		return "txt"
	}

	return "invalid"
}
