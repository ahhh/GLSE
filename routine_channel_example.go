// Inspired via the golang book: https://www.golang-book.com/books/intro/10
package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  var c chan string = make(chan string)
  go pinger(c)
  go ponger(c)
  go printer(c)
  var input string
  fmt.Scanln(&input)
}
// To be run as an infinate child thread
func pinger(c chan string) {
  for i := 0; ; i++ {
    //amt := time.Duration(rand.Intn(250))
    //time.Sleep(time.Millisecond * amt)
    c <- fmt.Sprintf("(ping %d)", i)
  }
}
// Runs our threads at different speeds
func ponger(c chan string) {
  for i := 0; ; i++ {
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
    c <- fmt.Sprintf("(pong %d)", i)
  }
}
// Prints our results at a constant interval
func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}
