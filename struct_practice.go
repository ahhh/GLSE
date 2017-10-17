package main

import (
  "fmt"
)

type newNew struct {
  Name string
  X int
}

func (p newNew) lolol() {
  p.X = (p.X + 2)
  fmt.Printf("Address of %s: %x\n", p.Name, p.Name)
}

func (p *newNew) whoa() {
  p.X = (p.X + 2)
}

func main(){
  lel := newNew{"Ez", 3}
  lol := newNew{"Ok", 5}
  fmt.Printf("Address of %s: %x\n", lel.Name, lel.Name)
  lel.lolol()
  fmt.Printf("Value of %s: %d\n", lol.Name, lol.X)
  lol.whoa()
  fmt.Printf("Value of %x: %d\n", lol.Name, lol.X)
  lol.lolol()
  fmt.Printf("Value of %x: %d\n", lol.Name, lol.X)
}
