package main

import (
  "os"
  //fp "path/filepath"

  "github.com/zsoltk-iw/guava/cmd"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
    os.Exit(1)
  }
}