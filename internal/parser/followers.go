package parser

import (
	"encoding/json"
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
	// TODO: Implement file reading
	// TODO: Implement JSON parsing
	// TODO: Extract usernames from string_list_data[0].value
	// TODO: Return map of usernames
	return nil, nil
}

// parseFollowersData parses the raw JSON data into FollowerEntry slice
func parseFollowersData(data []byte) ([]FollowerEntry, error) {
	// TODO: Implement JSON unmarshaling
	return nil, nil
}
