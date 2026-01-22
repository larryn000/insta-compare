package output

import (
	"io"

	"insta-compare/internal/compare"
)

// Format represents the output format type
type Format string

const (
	FormatText Format = "text"
	FormatJSON Format = "json"
)

// Writer handles output formatting and writing
type Writer struct {
	format Format
	writer io.Writer
}

// NewWriter creates a new output Writer
func NewWriter(w io.Writer, format Format) *Writer {
	return &Writer{
		format: format,
		writer: w,
	}
}

// Write outputs the comparison results in the configured format
func (w *Writer) Write(result *compare.CompareResult) error {
	// TODO: Check format type
	// TODO: Call appropriate formatter (text or JSON)
	// TODO: Write to output writer
	return nil
}

// writeText formats and writes results as human-readable text
func (w *Writer) writeText(result *compare.CompareResult) error {
	// TODO: Write header
	// TODO: Write numbered list of non-followers
	// TODO: Write total count
	return nil
}

// writeJSON formats and writes results as JSON
func (w *Writer) writeJSON(result *compare.CompareResult) error {
	// TODO: Marshal result to JSON
	// TODO: Write to output writer
	return nil
}
