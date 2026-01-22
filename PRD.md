# Product Requirements Document: Instagram Follow Comparison Service

## Overview

A command-line Golang service that compares Instagram followers and following data to identify users who are not following back.

## Problem Statement

Users want to identify accounts they follow on Instagram that do not follow them back. Instagram's data export provides separate JSON files for followers and following, but no built-in way to compare them.

## Input Data

### Following File (`following.json`)

```json
{
  "relationships_following": [
    {
      "title": "username",
      "string_list_data": [
        {
          "href": "https://www.instagram.com/_u/username",
          "timestamp": 1767643153
        }
      ]
    }
  ]
}
```

- **Structure**: Object with `relationships_following` array
- **Username location**: `title` field of each entry
- **Timestamp**: Unix timestamp of when the user was followed

### Followers File (`followers.json`)

```json
[
  {
    "title": "",
    "media_list_data": [],
    "string_list_data": [
      {
        "href": "https://www.instagram.com/username",
        "value": "username",
        "timestamp": 1767643173
      }
    ]
  }
]
```

- **Structure**: Array of objects
- **Username location**: `value` field within `string_list_data[0]`
- **Timestamp**: Unix timestamp of when the user followed

## Functional Requirements

### Core Features

1. **Parse JSON Files**
   - Read and parse `followers.json`
   - Read and parse `following.json`
   - Handle malformed JSON with appropriate error messages

2. **Extract Usernames**
   - Extract usernames from the `title` field in following data
   - Extract usernames from `string_list_data[0].value` in followers data

3. **Compare Lists**
   - Build a set of follower usernames for O(1) lookup
   - Iterate through following list and check against followers set
   - Identify users in following who are not in followers

4. **Output Results**
   - Display list of non-followers to stdout
   - Show total count of non-followers
   - Optionally output to a file

### Command-Line Interface

```bash
# Basic usage with default file names
insta-compare

# Specify custom file paths
insta-compare --followers ./followers.json --following ./following.json

# Output to file
insta-compare --output non-followers.txt

# JSON output format
insta-compare --format json
```

### CLI Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--followers` | `-f` | `followers.json` | Path to followers JSON file |
| `--following` | `-g` | `following.json` | Path to following JSON file |
| `--output` | `-o` | stdout | Output file path |
| `--format` | `-fmt` | `text` | Output format: `text` or `json` |

## Non-Functional Requirements

1. **Performance**: Process files with 10,000+ entries in under 1 second
2. **Memory**: Efficient memory usage with streaming where possible
3. **Error Handling**: Clear error messages for file not found, parse errors, etc.
4. **Cross-Platform**: Build for Linux, macOS, and Windows

## Technical Design

### Data Structures

```go
// FollowingFile represents the following.json structure
type FollowingFile struct {
    RelationshipsFollowing []FollowingEntry `json:"relationships_following"`
}

type FollowingEntry struct {
    Title          string           `json:"title"`
    StringListData []StringListItem `json:"string_list_data"`
}

// FollowerEntry represents each entry in followers.json
type FollowerEntry struct {
    Title          string           `json:"title"`
    MediaListData  []interface{}    `json:"media_list_data"`
    StringListData []StringListItem `json:"string_list_data"`
}

type StringListItem struct {
    Href      string `json:"href"`
    Value     string `json:"value,omitempty"`
    Timestamp int64  `json:"timestamp"`
}

// Result represents a non-follower
type Result struct {
    Username  string `json:"username"`
    URL       string `json:"url"`
    Timestamp int64  `json:"followed_at"`
}
```

### Algorithm

1. Load followers.json → Extract usernames → Store in `map[string]bool`
2. Load following.json → Iterate entries
3. For each following entry, check if username exists in followers map
4. If not found, add to results list
5. Output results

## Output Examples

### Text Format (default)

```
Users you follow who don't follow you back:
-------------------------------------------
1. tranxtruong
2. worldofxtra
3. wdnt.st
4. zohrankmamdani
5. jack_quaid

Total: 5 users
```

### JSON Format

```json
{
  "non_followers": [
    {
      "username": "tranxtruong",
      "url": "https://www.instagram.com/_u/tranxtruong",
      "followed_at": 1765503846
    }
  ],
  "total": 5,
  "following_count": 500,
  "followers_count": 495
}
```

## Project Structure

```
insta-compare/
├── main.go
├── go.mod
├── internal/
│   ├── parser/
│   │   ├── followers.go
│   │   └── following.go
│   ├── compare/
│   │   └── compare.go
│   └── output/
│       └── output.go
├── followers.json
├── following.json
└── README.md
```

## Success Criteria

- [ ] Successfully parses both JSON file formats
- [ ] Correctly identifies all non-followers
- [ ] CLI works with default and custom file paths
- [ ] Supports both text and JSON output formats
- [ ] Handles edge cases (empty files, missing fields)
- [ ] Includes unit tests for parser and comparison logic

## Future Enhancements (Out of Scope)

- Web UI interface
- Database storage for historical tracking
- Automatic unfollowing via Instagram API
- Mutual followers analysis
- Follow date analysis (how long ago you followed)
