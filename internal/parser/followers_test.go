package parser

import (
	"path/filepath"
	"testing"
)

func TestParseFollowers_ValidData(t *testing.T) {
	testFile := filepath.Join("testdata", "followers_valid.json")

	got, err := ParseFollowers(testFile)
	if err != nil {
		t.Fatalf("ParseFollowers() error = %v", err)
	}

	expected := map[string]bool{
		"user1": true,
		"user2": true,
		"user3": true,
	}

	if len(got) != len(expected) {
		t.Errorf("ParseFollowers() got %d entries, want %d", len(got), len(expected))
	}

	for username := range expected {
		if !got[username] {
			t.Errorf("ParseFollowers() missing expected username: %s", username)
		}
	}
}

func TestParseFollowers_EmptyArray(t *testing.T) {
	testFile := filepath.Join("testdata", "followers_empty.json")

	got, err := ParseFollowers(testFile)
	if err != nil {
		t.Fatalf("ParseFollowers() error = %v", err)
	}

	if len(got) != 0 {
		t.Errorf("ParseFollowers() got %d entries, want 0", len(got))
	}
}

func TestParseFollowers_InvalidJSON(t *testing.T) {
	testFile := filepath.Join("testdata", "followers_invalid.json")

	_, err := ParseFollowers(testFile)
	if err == nil {
		t.Error("ParseFollowers() expected error for invalid JSON, got nil")
	}
}

func TestParseFollowers_FileNotFound(t *testing.T) {
	_, err := ParseFollowers("testdata/nonexistent.json")
	if err == nil {
		t.Error("ParseFollowers() expected error for nonexistent file, got nil")
	}
}
