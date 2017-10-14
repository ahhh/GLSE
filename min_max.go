package main

import "fmt"

var max, min int

func main() {
  datArray := []int{1,2,3,4,5,6,7,8,9,0}

  datMax := maxz(datArray)
  fmt.Println("Max int in the array: ", datMax)

  datMin := minz(datArray)
  fmt.Println("Min int in the array: ", datMin)
}

func maxz(someArray []int) int {
  for _, num := range someArray {
    if num > max {
      max = num
    }
  }
  return max
}

func minz(someArray []int) int{
  for _, num := range someArray {
    if num < min {
      min = num
    }
  }
  return min
}
