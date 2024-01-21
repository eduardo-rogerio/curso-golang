package arquiterura

import (
  "runtime"
  "testing"
)

func TestDependente(t *testing.T) {
  t.Parallel()
  if runtime.GOARCH == "amd64" {
    t.Skip("Não funciona em arquiterura amd64")
  }
  t.Fail()

}
