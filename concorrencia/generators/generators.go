package main

import (
  "io/ioutil"
  "net/http"
  "regexp"
  "fmt"
)
// <-chan = canal somente leitura
func titulo(urls ...string) <-chan string {
  c := make(chan string)
  for _, url := range urls {
    go func(url string) {
      resp, _ := http.Get(url)
      html, _ := ioutil.ReadAll(resp.Body)

      r, _ := regexp.Compile("<title>(.*?)<\\/title>")
      c <- r.FindStringSubmatch(string(html))[1]
    }(url) 
  }
  return c
}
 
func main() {
  t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
  t2 := titulo("https://www.udemy.com", "https://www.youtube.com")

  fmt.Println("Primeiros:", <-t1,"|",<-t2)
  fmt.Println("Segundos:", <-t1,"|",<-t2)
}
