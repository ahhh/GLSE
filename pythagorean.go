package main

import (
  "math"
  "fmt"
  "os"
  "strconv"
)

func main(){

  fmt.Println("a^2 + b^2 = c^2")
  fmt.Println("Arg1 = a; Arg2 = b")

  leg1 := os.Args[1]
  leg2 := os.Args[2]
  var err error
  // Convert our command line strings to Ints
  var legA, legB int
  if legA, err = strconv.Atoi(leg1); err != nil {
    panic(err)
  }
  if legB, err = strconv.Atoi(leg2); err != nil {
    panic(err)
  }
  fmt.Println(int(legA), int(legB))
  hypotenuse := math.Sqrt(math.Pow(float64(legA),2) + math.Pow(float64(legB),2))
  fmt.Println(hypotenuse)

}
