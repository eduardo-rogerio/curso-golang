package main

import(
  "fmt"
  "strconv"
)

func main() {
  x := 2.4
  y := 2
  fmt.Println(x/float64(y))

  // se converter para inteiro ele pega somente a parte inteira do numero
  nota := 6.9
  notaFinal := int(nota)
  fmt.Println(notaFinal)

fmt.Println("Test " + strconv.Itoa(123))
}
