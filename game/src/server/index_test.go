package server

import(
  "testing"
  "gopkg.in/kataras/iris.v6"
  "gopkg.in/kataras/iris.v6/httptest"
)

func TestIndex(t *testing.T) {
  e := httptest.New(setupHttpServer(), t)
  e.GET("/").Expect().Status(iris.StatusOK)
}
