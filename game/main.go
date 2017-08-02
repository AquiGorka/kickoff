//+build !test

package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/AquiGorka/kickoff/game/server"
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
		Addr: ":" + os.Getenv("APP_PORT"),
		Handler: r,
	}
	log.Fatal(app.ListenAndServe())
}
