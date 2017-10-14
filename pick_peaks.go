package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	datArray := []int{7,1,2,7,9,43,2,5,7,9,0,3,4,1,4}
	myPeaks := PickPeaks(datArray)
	fmt.Println(myPeaks)
}

func PickPeaks(array []int) PosPeaks {
  var pos, peaks []int
  var tmp, large int
  peak_switch := false
  large_index := 100
  for index, item := range array {
    if index == 0 {
      tmp = item
    } else if item > tmp {
      large = item
      large_index = index
      peak_switch = true
    } else if item < tmp && peak_switch == true {
      pos = append(pos, large_index)
      peaks = append(peaks, large)
      peak_switch = false
    }
    tmp = item
  }
  if pos == nil {
    pos = make([]int, 0)
    peaks = make([]int, 0)
  }
	return PosPeaks{pos, peaks}
}
