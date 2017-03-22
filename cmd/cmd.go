// Package cmd provides a command-line interface for mmmgr.
package cmd

import (
	"io"
	"log"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "mmmgr",
	Short: "Manages multimedia",
	Long: `A command-line application and server to help you manage your
multimedia files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Usage(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}

// AddCommand adds subcommands to the root command
func AddCommand(cmds ...*cobra.Command) {
	root.AddCommand(cmds...)
}

// SetOutput sets the command's output
func SetOutput(w io.Writer) {
	root.SetOutput(w)
}

// SetArgs sets the args to parse; useful for testing.
func SetArgs(args []string) {
	root.SetArgs(args)
}
