package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/tigawanna/training/internals/database"
	"github.com/tigawanna/training/internals/utils"
	"github.com/uptrace/bun"
)

var (
	routes []string
	db     *bun.DB
)

func handle(mux *http.ServeMux, pattern string, handler http.HandlerFunc) {
	mux.HandleFunc(pattern, handler)
	routes = append(routes, pattern)
}

func BootstrapApi() {
	env := utils.GetEnv()

	var err error
	db, err = database.Open(env.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(context.Background()); err != nil {
		log.Fatal(err)
	}

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
