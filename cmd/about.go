package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

var (
  cmdAbout = &cobra.Command{
    Use: "about",
    Short: "describes the various use cases",
    Long: ``,
    RunE: runCmdAbout,
    SilenceUsage: true,
  }
)

func init() {
  RootCmd.AddCommand(cmdAbout)
}

func runCmdAbout(cmd *cobra.Command, args []string) error {
  usageMsg :=
`Help here ...
The application uses the 'romana/rlog' package for logging.
Please see https://github.com/romana/rlog for details on configuration
`
  fmt.Print(usageMsg)
  return nil
}