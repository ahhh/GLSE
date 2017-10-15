package main

import (
  "net/http"
  "io"
  "fmt"
  "time"
)
//Global var
var strCmd string

func main() {
  fmt.Println("Welcome friends!")
  strCmd = "defaultz"
  EZServer(":8080")
  secondz := 10
  fmt.Println(fmt.Sprintf("Serving content for %d seconds!", secondz))
  time.Sleep(1000*time.Duration(secondz)*time.Millisecond)
  strCmd = "something new!"
  fmt.Println(fmt.Sprintf("Serving different content for the %d seconds!", secondz))
  time.Sleep(1000*time.Duration(secondz)*time.Millisecond)
}

// Silly server defined on a specific port w/ certain paths and functions
func EZServer(port string) {
    go func() {
        http.HandleFunc("/", hello_world)
        http.HandleFunc("/cmd", cmd)
        err := http.ListenAndServe(port, nil)
        if err != nil {
          fmt.Println(err)
        }
    }()
}
// example handlers
func hello_world(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
func cmd(w http.ResponseWriter, r *http.Request ) {
	io.WriteString(w, strCmd)
}
