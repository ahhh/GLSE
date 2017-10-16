// GOOS=windows GOARCH=amd64 go build win_registry.go
package main

import (
  "fmt"
  reg "golang.org/x/sys/windows/registry"
  "os"
)

func helpMenu(){
  fmt.Println("Template code for messing w/ the Windows Registry")
  fmt.Println("Hidden Files in Explorer: win_reg.exe -hide")
  fmt.Println("Unhide Files in Explorer: win_reg.exe -unhide")
  fmt.Println("Set UserRunOnce Usage: win_reg.exe -set my_key C:\file_path")
  fmt.Println("Unset UserRunOnce Usage: ./win_reg.exe -unset my_key")
  fmt.Println("For help: ./win_reg.exe -h")
}

func main() {
  // First get our user Args
  if len(os.Args) > 1 {
    argFlag := os.Args[1]
    // Our Help Option
    if argFlag == "-h"{
      helpMenu()
    // Our Hide hidden files in Explorer option
    } else if os.Args[1] == "-hide"{
      ForceHiddenFiles()
      fmt.Println("Forced Hidden Files!")
    // Our Unhide hidden files in Explorer option
    } else if os.Args[1] == "-unhide" {
      ForceNoHiddenFiles()
      fmt.Println("Forced No Hidden Files in Explorer!")
    // Our persit to User RunOnce custom key option
    } else if os.Args[1] == "-set" {
      argKey := os.Args[2]
      argPath := os.Args[3]
      PersistUserRunOnce(argKey, argPath)
      fmt.Println("Set binary to persist to UserRunOnce: ", argPath)
    // Our unpersist a key option
    } else if os.Args[1] == "-unset" {
      argKey := os.Args[2]
      FixUserRunOnce(argKey)
      fmt.Println("Unset the persistance at: ", argKey)
    } else {
      helpMenu()
    }
  } else {
    helpMenu()
  }
}

// CreateKeyAndValue creates a new regestry key in a dynamic hive, dynamic path, dynamic object, and infers the type as either string or uint32 and creats the correct key type accordingly
func CreateKeyAndValue(hive reg.Key, keyPath string, keyObject string, keyValue interface{}) {
  // Create our key or see if it exists
  k, _, err := reg.CreateKey(hive, keyPath, reg.ALL_ACCESS)
  if err != nil {
    fmt.Println(err)
    return
  }
  // regardless if its new or created it we set it to our value
  defer k.Close()
  // switch on keyValue type to create different key type values
  switch v := keyValue.(type) {
    case string:
      keyValueSZ := keyValue.(string)
      k.SetStringValue(keyObject, keyValueSZ)
    case uint32:
      keyValueUI32 := keyValue.(uint32)
      k.SetDWordValue(keyObject, keyValueUI32)
    default:
      vstring := fmt.Sprintf("Info: %v", v)
      fmt.Println(vstring)
  }
}

// DeleteKeysValue deletes a dynamic keyobject in a dynamic hive and path
func DeleteKeysValue(hive reg.Key, keypath string, keyobject string) {
  k, err := reg.OpenKey(hive, keypath, reg.ALL_ACCESS)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer k.Close()
  k.DeleteValue(keyobject)
}

func ForceHiddenFiles() {
  var value uint32 = 0
  var value1 uint32 = 1
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\NOHIDDEN`, "CheckedValue", value)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\NOHIDDEN`, "DefaultValue", value)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SHOWALL`, "CheckedValue", value)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SHOWALL`, "DefaultValue", value)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SuperHidden`, "CheckedValue", value)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SuperHidden`, "DefaultValue", value)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Policies\Explorer`, "NoFolderOptions", value1)
}
func ForceNoHiddenFiles() {
  var value uint32 = 0
  var value1 uint32 = 1
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\NOHIDDEN`, "CheckedValue", value1)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\NOHIDDEN`, "DefaultValue", value1)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SHOWALL`, "CheckedValue", value1)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SHOWALL`, "DefaultValue", value1)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SuperHidden`, "CheckedValue", value1)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced\Folder\Hidden\SuperHidden`, "DefaultValue", value1)
  CreateKeyAndValue(reg.LOCAL_MACHINE, `Software\Microsoft\Windows\CurrentVersion\Policies\Explorer`, "NoFolderOptions", value)
}

// Persist to User RunOnceKey
func PersistUserRunOnce(myKey string, myExe string) {
  CreateKeyAndValue(reg.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, myKey, myExe)
}
func FixUserRunOnce(myKey string) {
  DeleteKeysValue(reg.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\RunOnce`, myKey)
}
