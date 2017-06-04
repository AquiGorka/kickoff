package server

import(
  "testing"
  "github.com/gavv/httpexpect"
  "net/http"
  "os"
)

func TestIndex(t *testing.T) {
  e := httpexpect.WithConfig(httpexpect.Config{
      BaseURL: "http://localhost:" + os.Getenv("APP_PORT"),
      Reporter: httpexpect.NewAssertReporter(t),
  })
  e.GET("/").
    Expect().
    Status(http.StatusOK).Body().Equal("")
  /*
  app := iris.New()
  app = HttpServer(app)
  e := httptest.New(app, t)
  e.GET("/").Expect().Status(iris.StatusOK).Body().Equal("")
  */
}
