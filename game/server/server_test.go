package server_test

import (
	"context"
	"github.com/AquiGorka/kickoff/game/server"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var app *http.Server
	go func() {
		// run the server
		r := mux.NewRouter()
		r = server.HTTPServer(r)
		r = server.WebsocketServer(r)
		app = &http.Server{
			Addr:    ":" + os.Getenv("APP_PORT"),
			Handler: r,
		}
		log.Fatal(app.ListenAndServe())
	}()
	// tests
	code := m.Run()
	// stop the server
	app.Shutdown(context.Background())
	// fin
	os.Exit(code)
}

func TestSocketServer(t *testing.T) {
	_, err := websocket.Dial("ws://localhost:"+os.Getenv("APP_PORT")+"/ws", "", "http://localhost/")
	if err != nil {
		t.Error("Error connecting to socket server: ", err)
	}
}

func TestHttpServer(t *testing.T) {
	_, err := http.Get("http://localhost:" + os.Getenv("APP_PORT") + "/ping")
	if err != nil {
		t.Error("Error connecting to http server: ", err)
	}
}
