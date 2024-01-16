package main
// Channel é uma forma de comunicação entre gorountines
// Channel é um tipo
import (
  "fmt"
  "time"
)
func doisTresQuatroVezes(base int, c chan int) {
  time.Sleep(time.Second)
  c <- 2 * base // enviando dados para o canal

  time.Sleep(time.Second)
  c <- 3 * base

  time.Sleep(time.Second)
  c <- 4 * base
}
func main() {
  c := make(chan int)
  go doisTresQuatroVezes(2,c)

  a, b := <-c,<-c
  fmt.Println(a, b)

  fmt.Println(<-c)
}
