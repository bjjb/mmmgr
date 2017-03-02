// Package main contains the main entry function for mmmgr.
package main

import (
	"fmt"
	"github.com/bjjb/mmmgr/files"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// The root command
var ui = &cobra.Command{
	Use:   "mmmgr",
	Short: "Manages multimedia",
	Long: `A command-line application and server to help you manage your
multimedia files.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var format string // template specified by flags

func init() {
	// Set up the configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.mmmgr")
	viper.SetEnvPrefix("mmmgr")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	findCommand.PersistentFlags().StringVarP(&format, "format", "f", "",
		"configure the command's output")

	ui.AddCommand(findCommand)
}

func main() {
	ui.Execute()
}

func makeExternalCommand(name string) (*cobra.Command, error) {
	debug("Making external command: %s", name)
	cmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Run the %s plugin", name),
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd, nil
}

var findCommand = &cobra.Command{
	Use:   "find PATH",
	Short: "Find and print multimedia files' metadata",
	Long: `
Scans the PATH for multimedia files, and prints the path name.
`,
	Run: func(cmd *cobra.Command, args []string) {
		outputter := func(f *files.File) {
			fmt.Println(f.AbsPath())
		}
		for _, arg := range args {
			switch {
			case files.IsDir(arg):
				for file := range files.Scan(arg) {
					outputter(file)
				}
			case files.IsMediaFile(arg):
				outputter(files.New(arg))
			}
		}
	},
}

func debug(fmt string, rest ...interface{}) {
	log.Printf(fmt, rest)
}
