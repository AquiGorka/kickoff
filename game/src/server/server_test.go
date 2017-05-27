package server

import (
  "testing"
  "os"
)

func TestMain(m *testing.M) {
  c := m.Run() // call all actual tests
  os.Exit(c)
}
