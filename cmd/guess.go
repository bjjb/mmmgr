package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/bjjb/mmmgr/guess"
)

// guessCmd represents the guess command
var guessCmd = &cobra.Command{
	Use:   "guess",
	Short: "Guess the type of a file",
	Long: `Tries to determine the multimedia properties of a given file.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range(args) {
			fmt.Printf("%v", guess.FromPath(path))
		}
	},
}

func init() {
	RootCmd.AddCommand(guessCmd)
}
