package main

import "fmt"

var count int

func main() {
  datArray := []string{"y","u", "lol"}
  // Test prints
  //fmt.Println(datArray[0])
  //fmt.Println(datArray[1])
  //fmt.Println(datArray[2])
  //fmt.Println(datArray)
  count = 0
  permute(0, datArray)
  fmt.Println(count)
}

// Our recursive permutation function
func permute(start int, substring []string) {
  // Main FOR loop for our factorial 
  for index := start; index < len(substring); index++ {
    // switch the two chars at the index
    temps := substring[start]
    substring[start] = substring[index]
    substring[index] = temps
    // recurse
    permute((start + 1), substring)
    // switch the chars back
    substring[index] = substring[start]
    substring[start] = temps
  }
  if (start == (len(substring)-1)) {
    // At the end of each permutation we print the line and a count
    fmt.Println(substring)
    count++
  }
}
