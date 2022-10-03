package config

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

func GetEnv() Env {
	return env
}
