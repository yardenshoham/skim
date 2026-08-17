module github.com/yardenshoham/skim

go 1.26.0

require (
	github.com/goccy/go-yaml v1.19.2
	github.com/spf13/cobra v1.10.2
	github.com/stretchr/testify v1.12.0
)

// https://github.com/goccy/go-yaml/pull/767
replace github.com/goccy/go-yaml => github.com/semihbkgr/go-yaml v0.0.0-20250623144847-1a8acaa8c12f

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
