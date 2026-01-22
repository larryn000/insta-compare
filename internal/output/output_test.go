package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"insta-compare/internal/compare"
)

func TestWriter_WriteText(t *testing.T) {
	result := &compare.CompareResult{
		NonFollowers: []compare.Result{
			{Username: "user1", URL: "https://instagram.com/_u/user1", Timestamp: 1000},
			{Username: "user2", URL: "https://instagram.com/_u/user2", Timestamp: 2000},
		},
		Total:          2,
		FollowingCount: 10,
		FollowersCount: 8,
	}

	var buf bytes.Buffer
	writer := NewWriter(&buf, FormatText)

	err := writer.Write(result)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	output := buf.String()

	if !strings.Contains(output, "Users you follow who don't follow you back:") {
		t.Error("Write() missing header")
	}

	if !strings.Contains(output, "1. user1") {
		t.Error("Write() missing user1")
	}

	if !strings.Contains(output, "2. user2") {
		t.Error("Write() missing user2")
	}

	if !strings.Contains(output, "Total: 2 users") {
		t.Error("Write() missing total count")
	}

	if !strings.Contains(output, "Following: 10 | Followers: 8") {
		t.Error("Write() missing following/followers counts")
	}
}

func TestWriter_WriteJSON(t *testing.T) {
	result := &compare.CompareResult{
		NonFollowers: []compare.Result{
			{Username: "user1", URL: "https://instagram.com/_u/user1", Timestamp: 1000},
		},
		Total:          1,
		FollowingCount: 5,
		FollowersCount: 4,
	}

	var buf bytes.Buffer
	writer := NewWriter(&buf, FormatJSON)

	err := writer.Write(result)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	var decoded compare.CompareResult
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Failed to decode JSON output: %v", err)
	}

	if decoded.Total != 1 {
		t.Errorf("JSON output Total = %d, want 1", decoded.Total)
	}

	if decoded.FollowingCount != 5 {
		t.Errorf("JSON output FollowingCount = %d, want 5", decoded.FollowingCount)
	}

	if decoded.FollowersCount != 4 {
		t.Errorf("JSON output FollowersCount = %d, want 4", decoded.FollowersCount)
	}

	if len(decoded.NonFollowers) != 1 {
		t.Fatalf("JSON output NonFollowers length = %d, want 1", len(decoded.NonFollowers))
	}

	if decoded.NonFollowers[0].Username != "user1" {
		t.Errorf("JSON output NonFollowers[0].Username = %s, want user1", decoded.NonFollowers[0].Username)
	}
}

func TestWriter_WriteText_Empty(t *testing.T) {
	result := &compare.CompareResult{
		NonFollowers:   []compare.Result{},
		Total:          0,
		FollowingCount: 5,
		FollowersCount: 5,
	}

	var buf bytes.Buffer
	writer := NewWriter(&buf, FormatText)

	err := writer.Write(result)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	output := buf.String()

	if !strings.Contains(output, "Total: 0 users") {
		t.Error("Write() missing zero total count")
	}
}

func TestWriter_DefaultFormat(t *testing.T) {
	result := &compare.CompareResult{
		NonFollowers:   []compare.Result{},
		Total:          0,
		FollowingCount: 0,
		FollowersCount: 0,
	}

	var buf bytes.Buffer
	writer := NewWriter(&buf, Format("unknown"))

	err := writer.Write(result)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	output := buf.String()

	// Should fall back to text format
	if !strings.Contains(output, "Users you follow") {
		t.Error("Write() should default to text format for unknown format")
	}
}
