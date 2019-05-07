package main

import (
	"fmt"
	"github.com/appscode/go/runtime"
	"os"
	"path/filepath"
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

	dir := os.Args[1]
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(runtime.GOPath(), dir)
	}
	fmt.Println("using repo:", dir)

}
