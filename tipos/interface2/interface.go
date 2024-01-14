package main

import "fmt"

type esportivo interface {
  ligarTurbo()
}

type ferrari struct {
  modelo string
  turboLigado bool
  velocidadeAtual int
}

func (f *ferrari) ligarTurbo()  {
  f.turboLigado = true
}

func main() {
  var carro esportivo = &ferrari{"F40", false, 0}
  carro.ligarTurbo()

  fmt.Println(carro)
}
