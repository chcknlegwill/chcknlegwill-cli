package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

// chcknlegwill-cli v1.0.1

func main() {
	// Define CLI flags ("-h", "--help", "-f")
	searchString := pflag.StringP("string", "f", "", "Search for a string in files and folders recursively.")
	help := pflag.BoolP("help", "h", false, "Show help message.")
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	// Show help if requested
	if *help || (pflag.NFlag() == 0) {
		fmt.Print("Usage:\nchcknlegwill-cli -f <string> | Search for a string within the current directory and folders recursively.\n")
		fmt.Print("chcknlegwill-cli -h --help Show this help message.\n")
		return
	}

	// Check if -f is provided but no value is given
	if *searchString == "" && pflag.NFlag() > 0 {
		fmt.Fprintf(os.Stderr, "Error: The -f flag requires a search string (e.g., -f keyword)\n")
		os.Exit(1)
	}

	// If -f is provided with a string, trigger search functionality
	if *searchString != "" {
		err := searchFiles(*searchString) // Pass the search string to searchFiles
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during search: %v\n", err)
			os.Exit(1)
		}
	}
}
