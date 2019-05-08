package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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

	sh := shell.NewSession()
	sh.SetEnv("GO111MODULE", "on")
	sh.SetDir(dir)
	sh.ShowCMD = true
	sh.PipeFail = true
	sh.PipeStdErrors = true

	gomodFile := filepath.Join(dir, "go.mod")
	if Exists(gomodFile) {
		data, err := ioutil.ReadFile("k8s_deps.json")
		if err != nil {
			log.Fatalln(err)
		}
		var kp []K8sPkg
		err = json.Unmarshal(data, &kp)
		if err != nil {
			log.Fatalln(err)
		}
		for _, x := range kp {
			if x.Repo != "" {
				continue
			}
			err = sh.Command("go", "get", "-u", fmt.Sprintf("%s@%s", x.Package, x.Version)).Run()
			if err != nil {
				fmt.Println(err)
				// continue
			}
		}
		return
	}

	glideFile := filepath.Join(dir, "glide.yaml")
	if Exists(glideFile) {
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

		err = sh.Command("go", "mod", "init").Run()
		if err != nil {
			fmt.Println(err)
		}
		for _, x := range cfg.Import {
			if x.Repo == "" {
				continue
			}

			repo := x.Repo
			repo = strings.ReplaceAll(repo, "https://", "")
			repo = strings.ReplaceAll(repo, "http://", "")
			repo = strings.ReplaceAll(repo, ".git", "")

			// go mod edit -replace=github.com/go-macaron/binding=github.com/tamalsaha/binding@pb
			err = sh.Command("go", "mod", "edit", fmt.Sprintf("-replace=%s=%s@%s", x.Package, repo, x.Version)).Run()
			if err != nil {
				fmt.Println(err)
				// continue
			}
			err = sh.Command("go", "mod", "tidy").Run()
			if err != nil {
				fmt.Println(err)
				// continue
			}
		}
		err = sh.Command("go", "mod", "tidy").Run()
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	depFile := filepath.Join(dir, "Gopkg.toml")
	if Exists(depFile) {
		fmt.Println("found Gopkg.toml: ", depFile)

		var cfg Dep
		if _, err := toml.DecodeFile(depFile, &cfg); err != nil {
			log.Fatalln(err)
		}

		err := sh.Command("go", "mod", "init").Run()
		if err != nil {
			fmt.Println(err)
		}
		for _, x := range append(cfg.Constraint, cfg.Override...) {
			if x.Source == "" {
				continue
			}

			repo := x.Source
			repo = strings.ReplaceAll(repo, "https://", "")
			repo = strings.ReplaceAll(repo, "http://", "")
			repo = strings.ReplaceAll(repo, ".git", "")

			tag := x.Version
			if x.Version == "" {
				if x.Revision == "" {
					tag = x.Branch
				} else {
					tag = x.Revision
				}
			}

			// go mod edit -replace=github.com/go-macaron/binding=github.com/tamalsaha/binding@pb
			err = sh.Command("go", "mod", "edit", fmt.Sprintf("-replace=%s=%s@%s", x.Name, repo, tag)).Run()
			if err != nil {
				fmt.Println(err)
				// continue
			}
			err = sh.Command("go", "mod", "tidy").Run()
			if err != nil {
				fmt.Println(err)
				// continue
			}
		}
		err = sh.Command("go", "mod", "tidy").Run()
		if err != nil {
			log.Fatalln(err)
		}
		return
	}
}
