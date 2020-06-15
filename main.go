/*
Copyright AppsCode Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	shell "github.com/codeskyblue/go-sh"
	"github.com/google/uuid"
	"github.com/hashicorp/go-getter"
	flag "github.com/spf13/pflag"
	"golang.org/x/mod/modfile"
)

var (
	sessionID      = uuid.New().String()
	desiredModFile = flag.String("desired-gomod", "", "Path of desired go.mod file (local file or url is accepted)")
)

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

	localfile := filepath.Join(os.TempDir(), sessionID, "go.mod")
	opts := func(c *getter.Client) error {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		c.Pwd = pwd
		return nil
	}
	err := getter.GetFile(localfile, *desiredModFile, opts)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(localfile)
	if err != nil {
		panic(err)
	}

	desiredMods, err := modfile.Parse(localfile, data, nil)
	if err != nil {
		panic(err)
	}

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
