package config

import "os"

func init() {
	e := os.Getenv("GO_ENV")
	env = Env{
		value: EnvValue(e),
	}
}
