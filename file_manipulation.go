package main

import (
  "os"
  "io/ioutil"
  "net/http"
  "fmt"
  "errors"
  )

func main(){

  // First get our user Args
  argFlag := os.Args[1]
  if argFlag == "-h"{
    fmt.Println("Download Usage: ./file_manipulation -d [url to download] [file path to save to]")
    fmt.Println("Copy Usage: ./file_manipulation -c [source file to copy] [file path to save to]")
    fmt.Println("For help: ./file_manipulation -h")
  } else if os.Args[1] == "-d"{
    argUrl := os.Args[2]
    argPath := os.Args[3]
    err := DownloadFile(argUrl, argPath)
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println("Succesfully downloaded file to: ", argPath)
    }
  } else if os.Args[1] == "-c" {
    argSource := os.Args[2]
    argDest := os.Args[3]
    err := CopyFile(argSource, argDest)
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println("Succesfully copied file from ", argSource, " to ", argDest)
    }
  }
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

func DownloadFile(url string, localPath string) error {
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  pageData, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  err = CreateFile(pageData, localPath)
  if err != nil {
    fmt.Println("Error creating your file")
    return err
  }
  return nil
}

func CopyFile(srcFile string, dstFile string) error {
  //Check if the file we are copying exists
  if Exists(srcFile) {
    // Read our file in as a byte array
    dat, err := ioutil.ReadFile(srcFile)
    if err != nil {
      return err
    }
    // Debug print the file contents
    //fmt.Print(string(dat))
    err = CreateFile(dat, dstFile)
    if err != nil {
      return err
    }
    // if no errors then we've created our file and can return no errors
    return nil
  } else {
    return errors.New("The srcFile does not exists")
  }

}
