package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "os"
  "errors"
)

func helpMenu(){
  fmt.Println("Download Usage: ./web_requester -d [url to download] [file path to save to]")
  fmt.Println("Get Page Usage: ./web_requester -g [url]")
  fmt.Println("For help: ./web_requester -h")
}

func main() {
  // First get our user Args
  if len(os.Args) > 1 {
    argFlag := os.Args[1]
    // Our Help Option
    if argFlag == "-h"{
      helpMenu()
    // Our Download Option
    } else if os.Args[1] == "-d"{
      argUrl := os.Args[2]
      argPath := os.Args[3]
      //fmt.Println(argUrl)
      //fmt.Println(argPath)
      err := DownloadFile(argUrl, argPath)
      if err != nil {
        fmt.Println(err)
      } else {
        fmt.Println("Succesfully downloaded file to: ", argPath)
      }
    // Our Page Fetch Option
    } else if os.Args[1] == "-g" {
      argUrl := os.Args[2]
      response := requestServer(argUrl)
      fmt.Println(response)
    }
  } else {
    helpMenu()
  }
}

func requestServer(url string) string {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
    return fmt.Sprintf("Error getting the url")
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
    return fmt.Sprintf("Error getting the page")
  }
  fmt.Printf("\nPage Content: \n`%s`", string(body))
  return "Success!"
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
