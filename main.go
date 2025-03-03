package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/ttacon/chalk"
)

// chcknlegwill-cli v1.0.3

func init() {

}

func main() {
	// Define CLI flags ("-h", "--help", "-f")
	searchString := pflag.StringP("string", "f", "", "Search for a string in files and folders recursively.")
	help := pflag.BoolP("help", "h", false, "Show this help message.")
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	//working - add a verbose flag that outputs skipped files
	//as well as optional thorough search that includes all files prefixed with "."
	//e.g. .DS_Store .gitignore etc...
	red := chalk.Red.NewStyle().WithForeground(chalk.Red)
	//fmt.Print(red.Style("Bruh moment"))

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
			fmt.Fprintf(os.Stderr, red.Style("Error during search: %v\n"), err)
			os.Exit(1)
		}
	}
}
