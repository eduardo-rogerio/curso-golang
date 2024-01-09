package main

import "fmt"

func main()  {
  i := 1

  // criando um ponteiro do tipo inteiro
  var p *int = nil

  p = &i // pegando o endereço de memoria da variavel i
  *p++ // pegando o valor da variavel i e fazendo um incremento
  i++

  fmt.Println(p, *p, i, &i)
}
