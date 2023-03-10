package main

import (
	"github.com/danielgtaylor/openapi-cli-generator/cli"
)

func main() {
	cli.Init(&cli.Config{
		AppName:   "orgs",
		EnvPrefix: "ORGS",
		Version:   "1.0.0",
	})

	atlasApiTransformedRegister(false)
	cli.Root.Execute()
}
