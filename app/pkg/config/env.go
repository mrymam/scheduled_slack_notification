package config

import "os"

type EnvValue string

const (
	EnvProd EnvValue = "prod"
	EnvDev  EnvValue = "dev"
	EnvTest EnvValue = "test"
)

var env Env

type Env struct {
	value EnvValue
}

func (e Env) IsProd() bool {
	return e.value == EnvProd
}

func (e Env) IsDev() bool {
	return e.value == EnvDev
}

func (e Env) IsTest() bool {
	return e.value == EnvTest
}

func GetEnv() Env {
	return env
}

func loadEnv() {
	e := os.Getenv("GO_ENV")
	env = Env{
		value: EnvValue(e),
	}
}
