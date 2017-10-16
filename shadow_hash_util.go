package main

import (
  "github.com/tredoe/osutil/user/crypt"
  "github.com/tredoe/osutil/user/crypt/apr1_crypt"
  "github.com/tredoe/osutil/user/crypt/md5_crypt"
  "github.com/tredoe/osutil/user/crypt/sha256_crypt"
  "github.com/tredoe/osutil/user/crypt/sha512_crypt"
  "os"
  "fmt"
)

func helpMenu(){
  fmt.Println("APR1 hash: ./shadow_hash_util -apr1 data_to_hash")
  fmt.Println("MD5 hash: ./shadow_hash_util -md5 data_to_hash")
  fmt.Println("SHA256 hash: ./shadow_hash_util -sha256 data_to_hash")
  fmt.Println("SHA512 hash: ./shadow_hash_util -sha512 data_to_hash")
  fmt.Println("For help: ./shadow_hash_util -h")
}

func main() {
  // First get our user Args
  if len(os.Args) > 1 {
    argFlag := os.Args[1]
    // Our Help Option
    if argFlag == "-h"{
      helpMenu()
    } else if os.Args[1] == "-apr1"{
      apr1, err := GenShadowHash(os.Args[2], "apr1")
      if err != nil {
        fmt.Println(err)
      } else {
        fmt.Println("APR1 Hash: ", apr1)
      }
    // Our Unhide hidden files in Explorer option
    } else if os.Args[1] == "-md5"{
      md5, err := GenShadowHash(os.Args[2], "md5")
      if err != nil {
        fmt.Println(err)
      } else {
        fmt.Println("MD5 Hash: ", md5)
      }
    // Our Unhide hidden files in Explorer option
    } else if os.Args[1] == "-sha256"{
      sha256, err := GenShadowHash(os.Args[2], "sha256")
      if err != nil {
        fmt.Println(err)
      } else {
        fmt.Println("SHA256 Hash: ", sha256)
      }
    // Our Unhide hidden files in Explorer option
    } else if os.Args[1] == "-sha512" {
      sha512, err := GenShadowHash(os.Args[2], "sha512")
      if err != nil {
        fmt.Println(err)
      } else {
        fmt.Println("SHA512 Hash: ", sha512)
      }
    } else {
      helpMenu()
    }
  } else {
    helpMenu()
  }
}

//Generate different Shadow file hash types
func GenShadowHash(input, hashType string) (string, error) {
  var box crypt.Crypter
  switch hashType {
  case "apr1":
    box = apr1_crypt.New()
  case "md5":
    box = md5_crypt.New()
  case "sha256":
    box = sha256_crypt.New()
  case "sha512":
    box = sha512_crypt.New()
  default:
    box = sha512_crypt.New()
  }
  hash, err := box.Generate([]byte(input), []byte{})
  if err != nil {
    return "", err
  }
  return hash, nil
}
