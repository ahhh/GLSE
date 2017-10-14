package main

import (
"io/ioutil"
"os"
"fmt"
"errors"
)

func main() {
  var a, b, z []byte
  argFlag := os.Args[1]
  // if the first flag is for help
  if argFlag == "-h"{
    fmt.Println("XOR Usage: ./xor -f first_file second_file output_file")
    fmt.Println("XOR Usage: ./xor -s first_string second_string")
    fmt.Println("For help: ./xor -h")
  // if the first flag is for file XORs
  } else if argFlag == "-f" {
    argFile1 := os.Args[2]
    argFile2 := os.Args[3]
    argOutput := os.Args[4]
    err := XorFiles(argFile1, argFile2, argOutput)
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println("File Xor Complete!")
    }
  // if the first flag is for string XORs
  } else if argFlag == "-s" {
    argString1 := os.Args[2]
    argString2 := os.Args[3]
    //fmt.Println(argString1, argString2)
    a = []byte(argString1)
    b = []byte(argString2)
    //fmt.Println(a[:], b[:])
    z = XorBytes(a[:], b[:])
    outString := string(z[:])
    fmt.Println(z[:])
    fmt.Println(outString)
  }
}

func XorFiles(file1 string, file2 string, outPut string) error {
  dat1, err := ioutil.ReadFile(file1)
  if err != nil{
    return err
  }
  dat2, err := ioutil.ReadFile(file2)
  if err != nil{
    return err
  }
  dat3 := XorBytes(dat1[:], dat2[:])
  err = CreateFile(dat3[:], outPut)
  if err != nil{
    return err
  } else {
    fmt.Println("Succesfully XORd the files, saved output file to: ", outPut)
  }
  return nil
}

func XorBytes(a []byte, b []byte) []byte {
  n := len(a)
	if len(b) < n {
		n = len(b)
	}
  var byte_dst [20]byte
  //fmt.Println(a[:])
  //fmt.Println(b[:])
  //fmt.Println(len(a))
  //fmt.Println(len(b))
	for i := 0; i < n; i++ {
    //fmt.Println(i)
		byte_dst[i] = a[i] ^ b[i]
	}
	return byte_dst[:]
}

func CreateFile(bytes []byte, path string) error {
  // Check if the file already exists
  if Exists(path) {
    return errors.New("The file to create already exists so we won't overwite it")
  }
  // write the lines to the file
  err := ioutil.WriteFile(path, bytes, 0700)
  if err != nil {
    return err
  }
  return nil
}

func Exists(path string) bool {
  // Run stat on a file
  _, err := os.Stat(path)
  // If it runs fine the file exists
  if err == nil {
    return true
  }
  // If stat fails then the file does not exist
  return false
}
