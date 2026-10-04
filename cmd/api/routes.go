package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/tigawanna/training/internals/utils"
)

var routes []string

func handle(mux *http.ServeMux, pattern string, handler http.HandlerFunc) {
	mux.HandleFunc(pattern, handler)
	routes = append(routes, pattern)
}

func BootstrapApi() {
	env := utils.GetEnv()
	mux := http.NewServeMux()

	handle(mux, "GET /{$}", indexHandler)
	handle(mux, "GET /users", usersHandler)
	handle(mux, "POST /users", createUserHandler)

	ln, err := net.Listen("tcp", ":"+env.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server running on http://localhost:%v\n", env.Port)
	log.Fatal(http.Serve(ln, mux))
}
