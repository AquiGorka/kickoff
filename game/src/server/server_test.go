package server_test

import (
  "testing"
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/adaptors/httprouter"
  "gopkg.in/kataras/iris.v6/httptest"
  "game/server"
  xwebsocket "golang.org/x/net/websocket"
)

func TestHttpServer(t *testing.T) {
  app := iris.New()
  app = server.HttpServer(app)
  e := httptest.New(app, t)
  e.GET("/").Expect().Status(iris.StatusOK)
}

func TestSocketServer(t *testing.T) {
  app := iris.New()
  app.Adapt(httprouter.New())
  app = server.WebsocketServer(app)
  go app.Listen(":8080")
  _, err := xwebsocket.Dial("ws://localhost:8080/ws", "", "http://localhost/")
  if (err != nil) {
    t.Error("Error connecting to socket server: ", err)
  }
}
