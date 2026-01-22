package compare

import (
	"insta-compare/internal/parser"
)

// Result represents a user who is not following back
type Result struct {
	Username  string `json:"username"`
	URL       string `json:"url"`
	Timestamp int64  `json:"followed_at"`
}

// CompareResult holds the full comparison results
type CompareResult struct {
	NonFollowers    []Result `json:"non_followers"`
	Total           int      `json:"total"`
	FollowingCount  int      `json:"following_count"`
	FollowersCount  int      `json:"followers_count"`
}

// FindNonFollowers compares following list against followers
// Returns users who are in following but not in followers
func FindNonFollowers(following []parser.FollowingUser, followers map[string]bool) *CompareResult {
	var nonFollowers []Result

	for _, user := range following {
		if !followers[user.Username] {
			nonFollowers = append(nonFollowers, Result{
				Username:  user.Username,
				URL:       user.URL,
				Timestamp: user.Timestamp,
			})
		}
	}

	return &CompareResult{
		NonFollowers:   nonFollowers,
		Total:          len(nonFollowers),
		FollowingCount: len(following),
		FollowersCount: len(followers),
	}
}
