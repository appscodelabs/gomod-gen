package main

type Glide struct {
	Package string `json:"package"`
	Import  []struct {
		Package string `json:"package"`
		Version string `json:"version,omitempty"`
		Repo    string `json:"repo,omitempty"`
		Vcs     string `json:"vcs,omitempty"`
	} `json:"import"`
}
