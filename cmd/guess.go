package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/bjjb/mmmgr/guess"
)

// guessCmd represents the guess command
var guessCmd = &cobra.Command{
	Use:   "guess FILE",
	Short: "Guess the type of a file",
	Long: `Tries to determine the multimedia properties of a given file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No files specified")
			os.Exit(1)
		}
		for _, path := range(args) {
			fmt.Printf("%v", guess.FromPath(path))
		}
	},
}

func init() {
	RootCmd.AddCommand(guessCmd)
}
