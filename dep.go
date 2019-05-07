package main

type Dep struct {
	Constraint []struct {
		Branch  string `toml:"branch,omitempty"`
		Name    string `toml:"name"`
		Version string `toml:"version,omitempty"`
	} `toml:"constraint"`
	Override []struct {
		Name     string `toml:"name"`
		Version  string `toml:"version,omitempty"`
		Branch   string `toml:"branch,omitempty"`
		Source   string `toml:"source,omitempty"`
		Revision string `toml:"revision,omitempty"`
	} `toml:"override"`
	Prune struct {
		UnusedPackages bool `toml:"unused-packages"`
		GoTests        bool `toml:"go-tests"`
	} `toml:"prune"`
}
