//+build !test

package main

import (
	"github.com/AquiGorka/kickoff/game/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	//
	r := mux.NewRouter()
	// http server
	r = server.HTTPServer(r)
	// websocket server
	r = server.WebsocketServer(r)
	//
	app := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: r,
	}
	log.Fatal(app.ListenAndServe())
}
