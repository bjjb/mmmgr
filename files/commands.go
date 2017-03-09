package files

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/template"
)

func init() {
	// Set up the command flags
	FindCommand.PersistentFlags().StringVarP(
		&findTemplateContent,
		"template",
		"t",
		`{{.Path}}{{"\n"}}`,
		"configure the command's output",
	)
}

/*
FindTemplate controls the output of the FindCommand - each File is passed
through it and printed.
*/
var FindTemplate = template.New("FindCommand")

/*
findTemplateContent contains the default or user-defined template through
which to pass all files found during the find command.
*/
var findTemplateContent string

/*
FindCommand implements the find command, using files.Scan
*/
var FindCommand = &cobra.Command{
	Use:   "find PATH",
	Short: "Find multimedia files",
	Long: `
Scans the PATH for multimedia files, and prints the path name. You can control
the output by passing a template which will be parsed instead for each
multimedia file found.
`,
	Run: func(cmd *cobra.Command, args []string) {
		templ, err := FindTemplate.Parse(findTemplateContent)
		if err != nil {
			log.Fatal(err)
		}
		for _, arg := range args {
			switch {
			case IsDir(arg):
				for f := range Scan(arg) {
					if err := templ.Execute(os.Stdout, f); err != nil {
						log.Fatal(err)
					}
					fmt.Println()
				}
			case IsMediaFile(arg):
				if err := templ.Execute(os.Stdout, New(arg)); err != nil {
					log.Fatal(err)
				}
				fmt.Println()
			}
		}
	},
}
