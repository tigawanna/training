package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func BootstrapApi() {
	env := GetEnv()
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	ln, err := net.Listen("tcp", ":"+env.port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server running on http://localhost:%v\n", env.port)
	log.Fatal(http.Serve(ln, mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Hello world")
}

type Environment string

const (
	Development Environment = "DEV"
	Production  Environment = "PROD"
)

type Env struct {
	port       string
	enviroment Environment
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
	return Env{port, environment}
}
