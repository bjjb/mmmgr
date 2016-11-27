package cmd

import (
	"fmt"
	"github.com/bjjb/mmmgr/guess"
	"github.com/spf13/cobra"
)

// guessCmd represents the guess command
var guessCmd = &cobra.Command{
	Use:   "guess FILE",
	Short: "Guess the type of a file",
	Long:  `Tries to determine the multimedia properties of a given file.`,
	Run: func(cmd *cobra.Command, args []string) {
		for g := range guess.GuessAll(args) {
			output(g)
		}
	},
}

func output(x interface{}) {
	fmt.Println(x)
}

func init() {
	RootCmd.AddCommand(guessCmd)
	guessCmd.Flags().StringVar(&templ, "template", "", "output template")
}
