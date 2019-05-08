package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func sort_k8s_deps() {
	data, err := ioutil.ReadFile("k8s_deps.json")
	if err != nil {
		log.Fatalln(err)
	}
	var kp []K8sPkg
	err = json.Unmarshal(data, &kp)
	if err != nil {
		log.Fatalln(err)
	}
	sort.Slice(kp, func(i, j int) bool { return kp[i].Package < kp[j].Package })

	str, err := json.MarshalIndent(kp, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile("k8s_deps.json", str, 0755)
}
