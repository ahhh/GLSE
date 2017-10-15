/* Developed while practicing with https://learnxinyminutes.com/docs/go/ */
// simple program that plays with language basics
package main
// single import
import "fmt"
// group import
import (
  "math"
  "math/rand"
  "time"
  )

func main(){
  // defer lets us run this last
  defer fmt.Println("Finished Exection!")
  // Simple print
  fmt.Println("Hello", "Friend")
  // private functions are lowercase
  //privateFunc()
  // Public functions are captialized
  PublicFunc()
}

// create a private interface of one function
type stringPair interface {
  string() string
}
// create a struct named pair with two fields, ints named x and y.
type pair struct {
  x, y float64
}
// create a function that uses our struct,
func (p pair) string() string { // p is called the "receiver"
  return fmt.Sprintf("(%f, %f)", p.x, p.y)
}
func privateFunc() {
  // decalre a bunch of variables without initalizing them
  var ranA, ranB, ranC, ranD, ranE float64
  // assign values to already initalized vars
  ranA = random(50, 75)
  ranB = random(1,10)
  ranC = random(2,99)
  ranD = random(4, 8)
  ranE = random(1, 3)
  // declare and initalize at the same time
  ranFinA := (((math.Pow(ranA, ranE)*ranE*ranC)/ranD)/ranB)
  ranFinB := ((math.Pow(ranB,ranD)*ranC*ranE)/ranD)
  // declare and initalize a struct
  myPair := pair{ranFinA, ranFinB}
  fmt.Println(myPair.string())
}
func random(min, max int) float64 {
  rand.Seed(time.Now().Unix())
  return float64(rand.Intn(max - min) + min)
}

func PublicFunc() {
  fmt.Println("lololololololololol")
  privateFunc()
  fmt.Println("lololololololololol")
}
