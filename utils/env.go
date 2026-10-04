package utils

import "os"

type Environment string

const (
	Development Environment = "DEV"
	Production  Environment = "PROD"
)

type Env struct {
	Port       string
	Enviroment Environment
}

func GetEnv() Env {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	environment := Environment(os.Getenv("ENVIROMENT"))
	if environment == "" {
		environment = Development
	}
	return Env{Port: port, Enviroment: environment}
}
