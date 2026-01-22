package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

// FollowingFile represents the following.json structure
type FollowingFile struct {
	RelationshipsFollowing []FollowingEntry `json:"relationships_following"`
}

// FollowingEntry represents each entry in the relationships_following array
type FollowingEntry struct {
	Title          string           `json:"title"`
	StringListData []StringListItem `json:"string_list_data"`
}

// FollowingUser represents a user being followed with metadata
type FollowingUser struct {
	Username  string
	URL       string
	Timestamp int64
}

// ParseFollowing reads and parses the following JSON file
// Returns a slice of FollowingUser
func ParseFollowing(filePath string) ([]FollowingUser, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading following file: %w", err)
	}

	file, err := parseFollowingData(data)
	if err != nil {
		return nil, fmt.Errorf("parsing following data: %w", err)
	}

	users := make([]FollowingUser, 0, len(file.RelationshipsFollowing))
	for _, entry := range file.RelationshipsFollowing {
		user := FollowingUser{
			Username: entry.Title,
		}
		if len(entry.StringListData) > 0 {
			user.URL = entry.StringListData[0].Href
			user.Timestamp = entry.StringListData[0].Timestamp
		}
		users = append(users, user)
	}

	return users, nil
}

// parseFollowingData parses the raw JSON data into FollowingFile
func parseFollowingData(data []byte) (*FollowingFile, error) {
	var file FollowingFile
	if err := json.Unmarshal(data, &file); err != nil {
		return nil, fmt.Errorf("unmarshaling JSON: %w", err)
	}
	return &file, nil
}
