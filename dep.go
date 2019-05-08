package main

type DepPkg struct {
	Name     string `toml:"name"`
	Version  string `toml:"version,omitempty"`
	Branch   string `toml:"branch,omitempty"`
	Source   string `toml:"source,omitempty"`
	Revision string `toml:"revision,omitempty"`
}

type Dep struct {
	Constraint []DepPkg `toml:"constraint"`
	Override   []DepPkg `toml:"override"`
	Prune      struct {
		UnusedPackages bool `toml:"unused-packages"`
		GoTests        bool `toml:"go-tests"`
	} `toml:"prune"`
}
