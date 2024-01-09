package main

import "fmt"

func imprimirResultado(nota float64)  {
  if nota >= 7 {
    fmt.Println("Aprovada com nota", nota)
  }else{
    fmt.Println("Reprovada com nota", nota)
  }
}

func main()  {
  imprimirResultado(7.3)
  imprimirResultado(5.1)
}
