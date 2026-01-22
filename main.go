package main

import (
	"flag"
	"fmt"
	"os"

	"insta-compare/internal/compare"
	"insta-compare/internal/output"
	"insta-compare/internal/parser"
)

func main() {
	// Define CLI flags
	followersPath := flag.String("followers", "followers.json", "Path to followers JSON file")
	followingPath := flag.String("following", "following.json", "Path to following JSON file")
	outputPath := flag.String("output", "", "Output file path (default: stdout)")
	formatFlag := flag.String("format", "text", "Output format: text or json")

	// Short flag aliases
	flag.StringVar(followersPath, "f", "followers.json", "Path to followers JSON file (shorthand)")
	flag.StringVar(followingPath, "g", "following.json", "Path to following JSON file (shorthand)")
	flag.StringVar(outputPath, "o", "", "Output file path (shorthand)")
	flag.StringVar(formatFlag, "fmt", "text", "Output format (shorthand)")

	flag.Parse()

	// Run the comparison
	if err := run(*followersPath, *followingPath, *outputPath, *formatFlag); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(followersPath, followingPath, outputPath, format string) error {
	// Parse followers file
	followers, err := parser.ParseFollowers(followersPath)
	if err != nil {
		return fmt.Errorf("failed to parse followers: %w", err)
	}

	// Parse following file
	following, err := parser.ParseFollowing(followingPath)
	if err != nil {
		return fmt.Errorf("failed to parse following: %w", err)
	}

	// Compare lists
	result := compare.FindNonFollowers(following, followers)

	// Setup output writer
	var out *os.File
	if outputPath != "" {
		out, err = os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer out.Close()
	} else {
		out = os.Stdout
	}

	writer := output.NewWriter(out, output.Format(format))

	// Write results
	if err := writer.Write(result); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}
