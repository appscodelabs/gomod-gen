package main

import (
	"fmt"
	"os"
)

// Example:
// gomod-tools <repo_path>
// gomod-tools github.com/appscode/voyager

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`Incorrect usage. Example of correct usage:
gomod-tools <repo_path>
gomod-tools github.com/appscode/voyager
`)
		os.Exit(1)
	}

}
