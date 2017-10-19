package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
  // Plan is to search a set of log files 
  // For requests that response size is greater than 1024 
  // HTTP method is POST

  pathToDir := os.Args[1]
  searchDirsForFiles(pathToDir)
}


func searchDirsForFiles(pathToDir string) {
  files, err := ioutil.ReadDir(pathToDir)
  if err != nil {
    fmt.Println(err)
  }
  for _, file := range files {
      fmt.Println(file.Name())
      if file.IsDir() {
        fmt.Println("File is a dir, not searching")
      } else {
        searchFilesForCriteria(pathToDir, file.Name())
      }
   }
}

// to be used on NGinx http logs
func searchFilesForCriteria(pathToDir, fileName string) {
  // first we need to read our file in
  criteria1, criteria2 := false, false
  // Recreate our full file path to read the files being searched
  fullPath := strings.Join([]string{pathToDir, filename}, "")
  fmt.Println(fullPath)
  fileData, err := ioutil.ReadFile(fullPath)
  if err !=nil {
    fmt.Println(err)
  }
  fileLines := strings.Split(string(fileData), "\n")
  for i, line := range fileLines {
    //fmt.Println(string(content))
    //criteria1 := strings.Contains(content, "")
    vals := strings.Split(string(line), " ")
    // Check our methods in the logs for POST
    for _, value := range vals {
      fmt.Println(string(value))
    }
    if string(vals[5]) == "POST" {
      criteria1 = true
    }
    //sets up criteria2
    respTime, err := strconv.Atoi(vals[10])
    if err != nil {
      fmt.Println(err)
    }
    if respTime > 1024 {
      criteria2 = true
    }
    if ((criteria1 && criteria2) == true) {
      fmt.Printf("The file %s, matched on line %d\n", fileName, i )
    }

  }
}
