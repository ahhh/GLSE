package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
    "math/rand"
    "time"
)

var size int = 100
var half int = size/2
var img = image.NewRGBA(image.Rect(0, 0, size, size))
var col color.Color

func main() {
  for i := 0; i < size; i++ {
    // Make a cross hatch pattern, where we layer horizontal lines, then vertical lines
    col = color.RGBA{uint8(random(100,200)), uint8(random(50,200)), uint8(random(0,100)), uint8(random(200,255))} // random1
    HLine(0, i, size)
    col = color.RGBA{uint8(random(0,50)), uint8(random(200,255)), uint8(random(100,200)), uint8(random(200,255))} // random2
    VLine(i, 0, size)
    // Draw random rectangles on top of our Cross hatch
    for e := i+1; e < size; e++ {
      col = color.RGBA{uint8(random(0,255)), uint8(random(0,255)), uint8(random(0,255)), uint8(random(200,255))} // random1
      Rect(random(0,half),random(0,half),random(0,half),random(0,half))
    }
  }
  f, err := os.Create("draw.png")
  if err != nil {
      panic(err)
  }
  defer f.Close()
  png.Encode(f, img)
}

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
  for ; x1 <= x2; x1++ {
    img.Set(x1, y, col)
  }
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
  for ; y1 <= y2; y1++ {
    img.Set(x, y1, col)
  }
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
  HLine(x1, y1, x2)
  HLine(x1, y2, x2)
  VLine(x1, y1, y2)
  VLine(x2, y1, y2)
}

func random(min, max int) int {
  rand.Seed(time.Now().Unix())
  return rand.Intn(max - min) + min
}
