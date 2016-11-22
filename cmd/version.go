package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print mmmgr version",
	Long:  `Print the version of mmmgr and exit`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("verbose") {
			fmt.Print("mmmgr v")
		}
		fmt.Println("0.0.1")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
