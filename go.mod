module github.com/yardenshoham/skim

go 1.26.0

require (
	github.com/goccy/go-yaml v1.19.2
	github.com/spf13/cobra v1.10.2
	github.com/stretchr/testify v1.12.1
)

// https://github.com/goccy/go-yaml/pull/767
replace github.com/goccy/go-yaml => github.com/semihbkgr/go-yaml v0.0.0-20250204052525-81dbf2619c42

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
)
