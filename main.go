/*
Package main contains the main entry function for mmmgr. The CLI is based on
http://github.com/spf13/cobra.
*/
package main

import (
	"github.com/bjjb/mmmgr/files"
	"github.com/bjjb/mmmgr/tvdb"
	"github.com/spf13/cobra"
	"log"
)

var format string // template specified by flags

func init() {
	// Set up the root command
	UI.AddCommand(files.FindCommand, tvdb.Command)
}

func main() {
	if err := UI.Execute(); err != nil {
		log.Fatalf("Error executing UI: %v", err)
	}
}

/*
UI is the root command.
*/
var UI = &cobra.Command{
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
