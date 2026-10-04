package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/tigawanna/training/internals/utils"
)

func BootstrapApi() {
	env := utils.GetEnv()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", indexHandler)
	mux.HandleFunc("POST /users", createUserHandler)

	ln, err := net.Listen("tcp", ":"+env.Port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server running on http://localhost:%v\n", env.Port)
	log.Fatal(http.Serve(ln, mux))
}
