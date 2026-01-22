package parser

import (
	"path/filepath"
	"testing"
)

func TestParseFollowing_ValidData(t *testing.T) {
	testFile := filepath.Join("testdata", "following_valid.json")

	got, err := ParseFollowing(testFile)
	if err != nil {
		t.Fatalf("ParseFollowing() error = %v", err)
	}

	if len(got) != 3 {
		t.Fatalf("ParseFollowing() got %d entries, want 3", len(got))
	}

	expected := []struct {
		username  string
		url       string
		timestamp int64
	}{
		{"user1", "https://www.instagram.com/_u/user1", 1767643153},
		{"user2", "https://www.instagram.com/_u/user2", 1767471834},
		{"user4", "https://www.instagram.com/_u/user4", 1766092048},
	}

	for i, exp := range expected {
		if got[i].Username != exp.username {
			t.Errorf("ParseFollowing()[%d].Username = %s, want %s", i, got[i].Username, exp.username)
		}
		if got[i].URL != exp.url {
			t.Errorf("ParseFollowing()[%d].URL = %s, want %s", i, got[i].URL, exp.url)
		}
		if got[i].Timestamp != exp.timestamp {
			t.Errorf("ParseFollowing()[%d].Timestamp = %d, want %d", i, got[i].Timestamp, exp.timestamp)
		}
	}
}

func TestParseFollowing_EmptyArray(t *testing.T) {
	testFile := filepath.Join("testdata", "following_empty.json")

	got, err := ParseFollowing(testFile)
	if err != nil {
		t.Fatalf("ParseFollowing() error = %v", err)
	}

	if len(got) != 0 {
		t.Errorf("ParseFollowing() got %d entries, want 0", len(got))
	}
}

func TestParseFollowing_InvalidJSON(t *testing.T) {
	testFile := filepath.Join("testdata", "following_invalid.json")

	_, err := ParseFollowing(testFile)
	if err == nil {
		t.Error("ParseFollowing() expected error for invalid JSON, got nil")
	}
}

func TestParseFollowing_FileNotFound(t *testing.T) {
	_, err := ParseFollowing("testdata/nonexistent.json")
	if err == nil {
		t.Error("ParseFollowing() expected error for nonexistent file, got nil")
	}
}
