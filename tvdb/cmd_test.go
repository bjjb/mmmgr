package tvdb

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCmd(t *testing.T) {
	commands := []*cobra.Command{
		rootCommand,
		languagesCommand,
		searchCommand,
	}
	for _, cmd := range commands {
		cmd.Run(&cobra.Command{Use: "x", Short: "", Long: ""}, []string{})
	}
}
