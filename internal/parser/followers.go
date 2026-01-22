package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

// FollowerEntry represents each entry in followers.json
type FollowerEntry struct {
	Title          string           `json:"title"`
	MediaListData  []interface{}    `json:"media_list_data"`
	StringListData []StringListItem `json:"string_list_data"`
}

// StringListItem represents the string_list_data structure
type StringListItem struct {
	Href      string `json:"href"`
	Value     string `json:"value,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

// ParseFollowers reads and parses the followers JSON file
// Returns a map of usernames for O(1) lookup
func ParseFollowers(filePath string) (map[string]bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading followers file: %w", err)
	}

	entries, err := parseFollowersData(data)
	if err != nil {
		return nil, fmt.Errorf("parsing followers data: %w", err)
	}

	followers := make(map[string]bool, len(entries))
	for _, entry := range entries {
		if len(entry.StringListData) > 0 && entry.StringListData[0].Value != "" {
			followers[entry.StringListData[0].Value] = true
		}
	}

	return followers, nil
}

// parseFollowersData parses the raw JSON data into FollowerEntry slice
func parseFollowersData(data []byte) ([]FollowerEntry, error) {
	var entries []FollowerEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("unmarshaling JSON: %w", err)
	}
	return entries, nil
}
