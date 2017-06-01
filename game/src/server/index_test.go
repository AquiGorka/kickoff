package server

import(
  "testing"
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/httptest"
)

func TestIndex(t *testing.T) {
  app := iris.New()
  HttpServer(app)
  e := httptest.New(app, t)
  e.GET("/").Expect().Status(iris.StatusOK).Body().Equal("")
}
