package cmd

import (
	"fmt"
	"github.com/bjjb/mmmgr/files"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan path",
	Short: "Scan the given path for multimedia",
	Long: `Scan the given path for multimedia files, printing out matches.
For every media file encountered, it simply prints the absolute path. Useful
for chaining commands together in shell scripts. For example:

$ for f in $(mmmgr scan .); do mmmgr add $f; done

The snippet above will add each file to the mmmgr library.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range args {
			filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				file := files.New(path)
				if err == nil && !info.IsDir() && file.MediaType != "" {
					fmt.Printf("%q\n", file.Path)
				}
				return err
			})
		}
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
}
