package main

import (
	"io/ioutil"
	"os"
	"path"
)

func writeTempFile(name string, fileData []byte) string {
	fileGuy, err := ioutil.TempFile("", name)
	if err != nil {
		return "Error creating temp file"
	}
	_, err = fileGuy.Write(fileData)
	if err != nil {
		return "Error writing the temp file"
	}
	err = fileGuy.Close()
	if err != nil {
		return "Error closing the temp file"
	}
	filepath := path.Join(fileGuy.Name())
	return filepath
}

func main() {
	// First get file to copy
	argFile := os.Args[1]
	dat, err := ioutil.ReadFile(argFile)
	if err != nil {
		panic(err)
	}
	location := writeTempFile("test", dat)
	println(location)
}
