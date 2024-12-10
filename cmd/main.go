package main

import (
  "fmt"
  "os"

  "github.com/akamensky/argparse"
)

func main() {
  parser := argparse.NewParser("WinBreak", "Description...")

  targetPtr := parser.String("t", "target", &argparse.Options{Required: true, Help: "Target machine IP"})

  err := parser.Parse(os.Args)
  if err != nil {
    fmt.Print(parser.Usage(err))
    return
  }

  target := *targetPtr

  fmt.Printf("Attacking %s\n", target)
}
