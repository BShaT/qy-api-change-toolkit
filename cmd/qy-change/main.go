package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/BShaT/qy-api-change-toolkit/internal/diff"
	"github.com/BShaT/qy-api-change-toolkit/internal/report"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	command := os.Args[1]
	if command == "help" || command == "--help" || command == "-h" {
		usage()
		return
	}
	if len(os.Args) < 4 {
		usage()
		os.Exit(2)
	}

	format := "json"
	flags := flag.NewFlagSet(command, flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	flags.StringVar(&format, "format", "json", "output format: json or markdown")
	if err := flags.Parse(os.Args[3:]); err != nil {
		os.Exit(2)
	}
	if flags.NArg() != 0 {
		usage()
		os.Exit(2)
	}

	oldPath, newPath := os.Args[2], os.Args[3]
	var changes []diff.Change
	var err error

	switch command {
	case "openapi":
		changes, err = compareDocuments(oldPath, newPath, diff.OpenAPI)
	case "schema":
		changes, err = compareDocuments(oldPath, newPath, diff.Schema)
	case "sitemap":
		changes, err = compareLoaded(oldPath, newPath, diff.LoadSitemap, diff.Sitemap)
	case "feed":
		changes, err = compareLoaded(oldPath, newPath, diff.LoadFeed, diff.Feed)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", command)
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var output []byte
	switch strings.ToLower(format) {
	case "json":
		output, err = report.JSON(changes)
	case "markdown", "md":
		output = report.Markdown(changes)
	default:
		fmt.Fprintf(os.Stderr, "unsupported format: %s\n", format)
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}

func compareDocuments(oldPath, newPath string, compare func(map[string]any, map[string]any) []diff.Change) ([]diff.Change, error) {
	old, err := diff.LoadDocument(oldPath)
	if err != nil {
		return nil, fmt.Errorf("load old document: %w", err)
	}
	new, err := diff.LoadDocument(newPath)
	if err != nil {
		return nil, fmt.Errorf("load new document: %w", err)
	}
	return compare(old, new), nil
}

func compareLoaded(oldPath, newPath string, load func(string) (map[string]any, error), compare func(map[string]any, map[string]any) []diff.Change) ([]diff.Change, error) {
	old, err := load(oldPath)
	if err != nil {
		return nil, fmt.Errorf("load old document: %w", err)
	}
	new, err := load(newPath)
	if err != nil {
		return nil, fmt.Errorf("load new document: %w", err)
	}
	return compare(old, new), nil
}

func usage() {
	fmt.Fprintln(os.Stderr, `Qy API Change Toolkit

Usage:
  qy-change openapi <old> <new> [--format json|markdown]
  qy-change schema <old> <new> [--format json|markdown]
  qy-change sitemap <old> <new> [--format json|markdown]
  qy-change feed <old> <new> [--format json|markdown]`)
}
