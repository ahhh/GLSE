// GOOS=windows GOARCH=amd64 go build win_powershell.go
package main

import (
  ps "github.com/gorillalabs/go-powershell"
  "github.com/gorillalabs/go-powershell/backend"
  //"golang.org/x/text/encoding/unicode"
  //"encoding/base64"
  "fmt"
  "time"
)

// https://github.com/jteeuwen/go-bindata
// We use items embeded in here from go-bindata assets shared to the target package
// Make sure to regenerate the included assets if you adding anything like so:
// go-bindata -nomemcopy -nometadata -pkg main -o bindata.go scripts/target_script.ps1

func main() {
  //fmt.Println("Running an embeded test script")
  //Select_Embedded_Script()
  fmt.Println("Clearing some event logs in 10sec!")
  time.Sleep(1000*time.Duration(10)*time.Millisecond)
  ClearEventLogs()
}

// ClearEventLogs : comment information goes here
func ClearEventLogs() {
  //Clear-EventLog Security, Application, Sysmon, System, "Windows PowerShell"
  RunPowerShell(`Clear-EventLog Security, Application, System`)
  RunPowerShell(`Clear-EventLog "Windows PowerShell"`)
  RunPowerShell(`Clear-EventLog Sysmon`)
}

// Run Target scripts
//Run our embeded powershell scripts here
//func Select_Embedded_Script() {
//  RunPowerShellScript("scripts/target_script.ps1")
//}

// RunPowerShell Gives a full powershell env
func RunPowerShell(cmd string) {
  back := &backend.Local{}
  shell, err := ps.New(back)
  if err != nil {
    fmt.Println(err)
  }
  defer shell.Exit()
  stdout, _, err := shell.Execute(cmd)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(stdout)
}

// RunPowerShellScript Function to run stored / encoded powershell scripts
//func RunPowerShellScript(scriptName string) {
//  data, err := Asset(scriptName)
//  if err != nil {
//    fmt.Println(err)
//  }
//  b64ScriptString, err := newEncodedPSScript(string(data))
//  if err != nil {
//    fmt.Println(err)
//  } else {
//    psCommand := fmt.Sprintf("powershell -Sta -NonInteractive -ExecutionPolicy bypass -EncodedCommand %s", b64ScriptString)
//    RunPowerShell(psCommand)
//  }
//}
//
//// newEncodedPSScript helper function for encoding powershell scripts to run
//func newEncodedPSScript(script string) (string, error) {
//  uni := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
//  encoded, err := uni.NewEncoder().String(script)
//  if err != nil {
//    return "", err
//  }
//  return base64.StdEncoding.EncodeToString([]byte(encoded)), nil
//}
