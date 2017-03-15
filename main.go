/*
Package main contains the main entry function for mmmgr. The CLI is based on
http://github.com/spf13/cobra.
*/
package main

import (
	_ "github.com/bjjb/mmmgr/cfg"
	"github.com/bjjb/mmmgr/cmd"
	_ "github.com/bjjb/mmmgr/files"
	_ "github.com/bjjb/mmmgr/musicbrainz"
	_ "github.com/bjjb/mmmgr/tmdb"
	_ "github.com/bjjb/mmmgr/tvdb"
)

func main() {
	cmd.Execute()
}
