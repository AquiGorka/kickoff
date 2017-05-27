package server

import (
  "gopkg.in/kataras/iris.v6"
)

func index(ctx *iris.Context) {
  ctx.HTML(iris.StatusOK, "")
}

