package compare

import (
	"testing"

	"insta-compare/internal/parser"
)

func TestFindNonFollowers_SomeNonFollowers(t *testing.T) {
	following := []parser.FollowingUser{
		{Username: "user1", URL: "https://instagram.com/_u/user1", Timestamp: 1000},
		{Username: "user2", URL: "https://instagram.com/_u/user2", Timestamp: 2000},
		{Username: "user3", URL: "https://instagram.com/_u/user3", Timestamp: 3000},
		{Username: "user4", URL: "https://instagram.com/_u/user4", Timestamp: 4000},
	}

	followers := map[string]bool{
		"user1": true,
		"user3": true,
	}

	result := FindNonFollowers(following, followers)

	if result.Total != 2 {
		t.Errorf("FindNonFollowers() Total = %d, want 2", result.Total)
	}

	if result.FollowingCount != 4 {
		t.Errorf("FindNonFollowers() FollowingCount = %d, want 4", result.FollowingCount)
	}

	if result.FollowersCount != 2 {
		t.Errorf("FindNonFollowers() FollowersCount = %d, want 2", result.FollowersCount)
	}

	expectedNonFollowers := map[string]bool{
		"user2": true,
		"user4": true,
	}

	for _, nf := range result.NonFollowers {
		if !expectedNonFollowers[nf.Username] {
			t.Errorf("FindNonFollowers() unexpected non-follower: %s", nf.Username)
		}
	}
}

func TestFindNonFollowers_AllFollowBack(t *testing.T) {
	following := []parser.FollowingUser{
		{Username: "user1", URL: "https://instagram.com/_u/user1", Timestamp: 1000},
		{Username: "user2", URL: "https://instagram.com/_u/user2", Timestamp: 2000},
	}

	followers := map[string]bool{
		"user1": true,
		"user2": true,
		"user3": true,
	}

	result := FindNonFollowers(following, followers)

	if result.Total != 0 {
		t.Errorf("FindNonFollowers() Total = %d, want 0", result.Total)
	}

	if len(result.NonFollowers) != 0 {
		t.Errorf("FindNonFollowers() NonFollowers length = %d, want 0", len(result.NonFollowers))
	}
}

func TestFindNonFollowers_NoneFollowBack(t *testing.T) {
	following := []parser.FollowingUser{
		{Username: "user1", URL: "https://instagram.com/_u/user1", Timestamp: 1000},
		{Username: "user2", URL: "https://instagram.com/_u/user2", Timestamp: 2000},
	}

	followers := map[string]bool{
		"user3": true,
		"user4": true,
	}

	result := FindNonFollowers(following, followers)

	if result.Total != 2 {
		t.Errorf("FindNonFollowers() Total = %d, want 2", result.Total)
	}
}

func TestFindNonFollowers_EmptyFollowing(t *testing.T) {
	following := []parser.FollowingUser{}
	followers := map[string]bool{"user1": true}

	result := FindNonFollowers(following, followers)

	if result.Total != 0 {
		t.Errorf("FindNonFollowers() Total = %d, want 0", result.Total)
	}

	if result.FollowingCount != 0 {
		t.Errorf("FindNonFollowers() FollowingCount = %d, want 0", result.FollowingCount)
	}
}

func TestFindNonFollowers_EmptyFollowers(t *testing.T) {
	following := []parser.FollowingUser{
		{Username: "user1", URL: "https://instagram.com/_u/user1", Timestamp: 1000},
	}
	followers := map[string]bool{}

	result := FindNonFollowers(following, followers)

	if result.Total != 1 {
		t.Errorf("FindNonFollowers() Total = %d, want 1", result.Total)
	}

	if result.FollowersCount != 0 {
		t.Errorf("FindNonFollowers() FollowersCount = %d, want 0", result.FollowersCount)
	}
}

func TestFindNonFollowers_PreservesMetadata(t *testing.T) {
	following := []parser.FollowingUser{
		{Username: "user1", URL: "https://instagram.com/_u/user1", Timestamp: 1234567890},
	}
	followers := map[string]bool{}

	result := FindNonFollowers(following, followers)

	if len(result.NonFollowers) != 1 {
		t.Fatalf("FindNonFollowers() expected 1 non-follower, got %d", len(result.NonFollowers))
	}

	nf := result.NonFollowers[0]
	if nf.Username != "user1" {
		t.Errorf("NonFollower.Username = %s, want user1", nf.Username)
	}
	if nf.URL != "https://instagram.com/_u/user1" {
		t.Errorf("NonFollower.URL = %s, want https://instagram.com/_u/user1", nf.URL)
	}
	if nf.Timestamp != 1234567890 {
		t.Errorf("NonFollower.Timestamp = %d, want 1234567890", nf.Timestamp)
	}
}
