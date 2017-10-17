package main

import (
  "os"
  "fmt"
  "strconv"
)

func main() {
  if len(os.Args) > 1 {
    arg1 := os.Args[1]
    arg2 := os.Args[2]
    arg3 := os.Args[3]
    var argi1, argi2, argi3 int
    argi1, err := strconv.Atoi(arg1)
    if err != nil {
      panic(err)
    }
    argi2, err = strconv.Atoi(arg2)
    if err != nil {
      panic(err)
    }
    argi3, err = strconv.Atoi(arg3)
    if err != nil {
      panic(err)
    }
    resultz := IsTriangle(argi1, argi2, argi3)
    if resultz == true {
      fmt.Println("The ints provided can form a triangle")
    } else {
      fmt.Println("The ints provided can not form a triangle")
    }
  } else {
    fmt.Println("Please provide 3 integers to see if they can form a triangle")
  }
}

func IsTriangle(a, b, c int) bool {
    var t1, t2, t3 bool

    if ((a + b) > c) { t1 = true
    } else { t1 = false }

    if ((a + c) > b) { t2 = true
    } else { t2 = false }

    if ((b + c) > a) { t3 = true
    } else { t3 = false }

    if (t1 && t2 && t3) == true { return true
    } else { return false }
}
