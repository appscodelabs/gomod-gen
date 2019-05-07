package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/appscode/go/runtime"
	shell "github.com/codeskyblue/go-sh"
	"sigs.k8s.io/yaml"
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
		dir = filepath.Join(runtime.GOPath(), "src", dir)
	}
	fmt.Println("using repo:", dir)

	glideFile := filepath.Join(dir, "glide.yaml")
	data, err := ioutil.ReadFile(glideFile)
	if err != nil {
		if os.IsNotExist(err) {
			// try for dep
		}
		log.Fatalln(err)
	}
	fmt.Println("found glide.yaml: ", glideFile)
	var cfg Glide
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	sh := shell.NewSession()
	sh.SetEnv("GO111MODULE", "on")
	sh.SetDir(dir)
	sh.ShowCMD = true
	sh.PipeFail = true
	sh.PipeStdErrors = true

	err = sh.Command("go", "mod", "init").Run()
	if err != nil {
		log.Fatalln(err)
	}
	for _, x := range cfg.Import {
		// go mod edit -replace=github.com/go-macaron/binding=github.com/tamalsaha/binding@pb
		err = sh.Command("go", "mod", "edit", fmt.Sprintf("-replace=%s=%s@%s", x.Package, x.Repo, x.Version)).Run()
		if err != nil {
			fmt.Println(err)
			// continue
		}
	}
	err = sh.Command("go", "mod", "tidy").Run()
	if err != nil {
		log.Fatalln(err)
	}
}
