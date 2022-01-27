package runtimeenvironment

import (
	"os"
	"strings"
)

const (
	LinkConverterEnvKey   = "LinkConverter_ENV"
	LinkConverterEnvLocal = "local"
)

var (
	Environment string
)

func loadVariables() {
	Environment = strings.TrimSpace(os.Getenv(LinkConverterEnvKey))

	if Environment == "" {
		Environment = LinkConverterEnvLocal
	}

}

func init() {
	loadVariables()
}
