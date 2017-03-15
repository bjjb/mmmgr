package main

import (
	"os"

	"github.com/bjjb/mmmgr/cmd"
)

func Example_scan() {
	cmd.SetOutput(os.Stdout)
	cmd.SetArgs([]string{"-h"})
	main()
	// Output:
	// A command-line application and server to help you manage your
	// multimedia files.
	//
	// Usage:
	//   mmmgr [flags]
	//   mmmgr [command]
	//
	// Available Commands:
	//   find        find and inspect multimedia files
	//   tvdb        interact with the TVDB
	//
	// Use "mmmgr [command] --help" for more information about a command.
}
