package main

import(
  "os"
  "os/exec"
  "fmt"
)

func helpMenu(){
  fmt.Println("Usage: ./run_cmd hostname -a")
}

func main() {
    // First get our user Args
  if len(os.Args) > 1 {
    argCmd := os.Args[1]
    resultz, err := RunCommand(argCmd, []string{})
    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println(resultz)
    }
  } else {
    helpMenu()
  }
}

func RunCommand(cmd string, args []string) (string, error) {
  out, err := exec.Command(cmd, args...).CombinedOutput()
  if err != nil {
    return string(out), err
  }
  return string(out), nil
}
