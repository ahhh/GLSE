package main

import "fmt"
import "time"
import "math/rand"

func main() {
  go i(0)
  for y := 0; y < 10; y++ {
    go i(y)
  }
  var input string
  fmt.Println("Enter a single string")
  // Scan for user input and collect it in the input var
  fmt.Scanln(&input)
  for z := 0; z < len(input); z++ {
    //go s_parts(input, z)
    go s_full(input, z)
  }
  time.Sleep(1000*time.Duration(5)*time.Millisecond)
  fmt.Println("Done!")
}

func s_full(in string, index int) {
  for i := 0; i < len(in); i++ {
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
    fmt.Println(in, ":", index, ":", i)
  }
}

func s_parts(in string, index int) {
  for i, part := range in {
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
    fmt.Println(part, ":", index, ":", i)
  }
}

func i(n int) {
  for x := 0; x < 10; x++ {
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
    fmt.Println(n, ":", x)
  }
}
