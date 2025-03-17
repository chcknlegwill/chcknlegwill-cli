package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

//Going to try and keep other imports to a minimum - only ones that
//I really need.

// chcknlegwill-cli v1.0.3

func init() {

}

func main() {
	// Define CLI flags ("-h", "--help", "-f")
	searchString := pflag.StringP("search", "f", "", "Search for a string in files and folders recursively.")
	list := pflag.BoolP("list", "l", false, "List the entire directory from the root (/).")
	help := pflag.BoolP("help", "h", false, "Show this help message.")
	pflag.CommandLine.SortFlags = false
	pflag.Parse()

	//working - add a verbose flag that outputs skipped files
	//as well as optional thorough search that includes all files prefixed with "."
	//e.g. .DS_Store .gitignore etc...
	//red := chalk.Red.NewStyle().WithForeground(chalk.Red)
	//fmt.Print(red.Style("Bruh moment"))

	// Show help if requested
	if *help || (pflag.NFlag() == 0) {
		fmt.Println("Usage:\nchcknlegwill-cli -f <string> | -l | -h | | --help")
		fmt.Println("-f, --search <string>    Search for a string in files and folders recursivley.")
		fmt.Println("-l, --list       	 List the entire directory from current directory.")
		fmt.Println("-h, --help	    	 Show this help message.")
		return
	}

	// Check if -f is provided but no value is given
	iFlag := pflag.Lookup("f")
	if iFlag != nil && iFlag.Changed && *searchString == "" {
		fmt.Fprintf(os.Stderr, "Error: The -f flag requires a search string (e.g., -f keyword)\n")
		os.Exit(1)
	}

	// If -f is provided with a string, trigger search functionality

	//TODO: Update with searching in specifiec directories e.g. if you are in one dir
	//and want to search a different folder you would write:
	//chcknlegwill-cli -f <string_to_search> <directory_to_search>
	if *searchString != "" {
		err := searchFiles(*searchString) // Pass the search string to searchFiles
		if err != nil {
			//fmt.Fprintf(os.Stderr, red.Style("Error during search: %v\n"), err)
			os.Exit(1)
		}
	}

	if *list {
		err := listDirectoryStructure(".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error listing directory structure: %v\n", err)
			os.Exit(1)
		}
		return
	}
}
