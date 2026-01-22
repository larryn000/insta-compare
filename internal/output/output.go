package output

import (
	"encoding/json"
	"fmt"
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
	switch w.format {
	case FormatJSON:
		return w.writeJSON(result)
	default:
		return w.writeText(result)
	}
}

// writeText formats and writes results as human-readable text
func (w *Writer) writeText(result *compare.CompareResult) error {
	fmt.Fprintln(w.writer, "Users you follow who don't follow you back:")
	fmt.Fprintln(w.writer, "-------------------------------------------")

	for i, user := range result.NonFollowers {
		fmt.Fprintf(w.writer, "%d. %s\n", i+1, user.Username)
	}

	fmt.Fprintln(w.writer)
	fmt.Fprintf(w.writer, "Total: %d users\n", result.Total)
	fmt.Fprintf(w.writer, "Following: %d | Followers: %d\n", result.FollowingCount, result.FollowersCount)

	return nil
}

// writeJSON formats and writes results as JSON
func (w *Writer) writeJSON(result *compare.CompareResult) error {
	encoder := json.NewEncoder(w.writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
