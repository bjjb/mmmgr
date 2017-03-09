package tvdb

import (
	"github.com/spf13/cobra"
	"log"
)

/*
Command is a cobra.Command which is to be added to the root command.
*/
var Command = &cobra.Command{
	Use:   "tvdb",
	Short: "interact with the TVDB",
	Long: `
A basic client for working with The TVDB (https://thetvdb.com).
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Usage(); err != nil {
			log.Fatal(err)
		}
	},
}
