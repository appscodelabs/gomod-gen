package main

type K8sPkg struct {
	Package string `json:"package"`
	Version string `json:"version"`
	Repo    string `json:"repo,omitempty"`
	Vcs     string `json:"vcs,omitempty"`
}
