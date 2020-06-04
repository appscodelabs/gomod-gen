package main

import (
	"fmt"
	"io/ioutil"
	"os"

	shell "github.com/codeskyblue/go-sh"
	flag "github.com/spf13/pflag"
	"golang.org/x/mod/modfile"
)

var desiredModFile = flag.String("desired-gomod", "", "Path of desired go.mod file")

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

	data, err := ioutil.ReadFile(*desiredModFile)
	if err != nil {
		panic(err)
	}

	desiredMods, err := modfile.ParseLax(*desiredModFile, data, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(desiredMods.Require)

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

	data, err = ioutil.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}
	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		panic(err)
	}

	for _, x := range desiredMods.Require {
		err = f.AddRequire(x.Mod.Path, x.Mod.Version)
		if err != nil {
			panic(err)
		}
	}

	for _, x := range desiredMods.Replace {
		err = f.DropReplace(x.Old.Path, x.Old.Version)
		if err != nil {
			panic(err)
		}
	}
	for _, x := range desiredMods.Replace {
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
}
