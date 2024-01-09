package main

import "fmt"

func main()  {
 var b byte = 3
  fmt.Println(b)

  i := 3 
  fmt.Println(i)

  // atribuição multipla
  x,y := 1,2
  fmt.Println(x,y)

  // trocando o valor de variaveis usando atribuição multipla
  x,y = y,x
  fmt.Println(x,y)
} 
