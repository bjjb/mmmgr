package files

import (
	"fmt"

	"github.com/bjjb/mmmgr/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(findCommand)
}

var findCommand = &cobra.Command{
	Use:   "find PATH",
	Short: "find and inspect multimedia files",
	Long: `
Scans the PATH for multimedia files, and prints the path name. You can control
the output by passing a template which will be parsed instead for each
multimedia file found.
		`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			switch {
			case IsDir(arg):
				for f := range Scan(arg) {
					fmt.Println(f)
				}
			case IsMediaFile(arg):
				fmt.Println(New(arg))
			}
		}
	},
}
