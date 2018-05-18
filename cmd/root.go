package cmd

import (
  // "os"
  // "fmt"
  // "strings"
  // fp "path/filepath"

  "github.com/spf13/cobra"
)

var (
  RootCmd = &cobra.Command{
    Use:                "guava",
    Short:              "GUAV server",
    PersistentPreRunE:  runCmdInit,
    SilenceErrors:      false,
    SilenceUsage:       false,
  }

  gsconf = GuavaConfig{}
)

func init() {
  // Global flags
  RootCmd.PersistentFlags().StringVarP(&gsconf.ServerPort, "port", "p", "", "The port the server should be listening")
}

func runCmdInit(cmd *cobra.Command, args []string) error {
  // Default is 80
  if !gsconf.validatePort() {
    gsconf.ServerPort = "8080"
  }

  return nil
}