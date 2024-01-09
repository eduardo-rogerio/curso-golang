package main

import (
  "math"
  "fmt"
)

func main()  {
  const PI float64 = 3.1415
  var raio = 3.2

  area := PI * math.Pow(raio,2)
  fmt.Println("A área da circunferência é", area)
}
