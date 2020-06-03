package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	shell "github.com/codeskyblue/go-sh"
	flag "github.com/spf13/pflag"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

var modJsonFile = flag.String("gomod-json-file", "", "Path of go.mod.json file")

type Module struct {
	Path    string
	Version string
}

type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
}

type Require struct {
	Path     string
	Version  string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}

// exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	flag.Parse()

	sh := shell.NewSession()
	sh.ShowCMD = true
	sh.PipeFail = true
	sh.PipeStdErrors = true

	if !exists("go.mod") {
		err := sh.Command("go", "mod", "init").Run()
		if err != nil {
			panic(err)
		}
	}

	data, err := ioutil.ReadFile(*modJsonFile)
	if err != nil {
		panic(err)
	}
	var kp GoMod
	err = json.Unmarshal(data, &kp)
	if err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}
	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		panic(err)
	}

	requires := make([]*modfile.Require, 0, len(kp.Require))
	for _, x := range kp.Require {
		requires = append(requires, &modfile.Require{
			Mod: module.Version{
				Path:    x.Path,
				Version: x.Version,
			},
		})
	}
	f.SetRequire(requires)
	for _, x := range kp.Replace {
		err = f.DropReplace(x.Old.Path, x.Old.Version)
		if err != nil {
			panic(err)
		}
	}
	for _, x := range kp.Replace {
		err = f.AddReplace(x.Old.Path, x.Old.Version, x.New.Path, x.New.Version)
		if err != nil {
			panic(err)
		}
	}

	data, err = f.Format()
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("go.mod", data, 0644)
	if err != nil {
		panic(err)
	}

	err = sh.Command("go", "mod", "tidy").Run()
	if err != nil {
		panic(err)
	}
	err = sh.Command("go", "mod", "vendor").Run()
	if err != nil {
		panic(err)
	}
}
