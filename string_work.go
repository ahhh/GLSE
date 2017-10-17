package main

import (
  "fmt"
  //"os"
  "strings"
)

func main() {
  var lol, lel string
  var lolol, lelel, ss []string
  lol = "ok I just fixed it"
  lel = "ok lets dance"
  lo := "we might have a problem"
  le := "what should we do"
  lolol = append(lolol, lo)
  lolol = append(lolol, le)
  lelel = append(lelel, lol)
  lelel = append(lelel, lel)
  ss = append(lolol, lelel[0])
  ss = append(ss, lelel[1])
  with_Slices(ss)
  with_Strings(ss)
}


func with_Strings(work []string) {
  fmt.Printf("dat work: %s\n", work)
  for i, sent :=  range work {
    fmt.Printf("sentance %d: %s\n", i, sent)
    wordz := strings.Split(sent, " ")
    for j, wordo := range wordz {
      fmt.Printf("word %d: %s\n", j, wordo)
      reado := strings.NewReader(wordo)
      for k, _ := range wordo {
        charo, _, _ := reado.ReadRune()
        fmt.Printf("char %d: %s\n", k, string(charo))
      }
    }
  }
}

func with_Slices(work []string) {
  fmt.Printf("dat work: %s\n", work)
  for i, sent := range work {
    fmt.Printf("sentance %d: %s\n", i, string(sent))
    for j, charz := range sent {
      fmt.Printf("char %d: %s\n", j, string(charz))
    }
  }
}
