package tvdb

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/bjjb/mmmgr/cmd"
	"github.com/spf13/cobra"
)

var jsonIndentation = "  "

var options = struct {
	search struct {
		params           bool
		imdbID, zap2itID string
	}
	json, trace bool
	template    string
}{}

var templates = map[string]string{
	"languages": `
{{if headers -}} {{printf "%-8s %-4s %s" "ID" "Code" "Name"}} {{- end}}
{{range . -}}
{{printf "%-8d %-4s %s" .ID .Abbr .Name}}
{{end}}`,
	"search": `
{{if headers -}} {{printf "%-8s %-10s %s" "ID" "Date" "Name"}} {{- end}}
{{range . -}}{{printf "%-8d %-10s %s" .ID .FirstAired .Name}}
{{end}}`,
}

func init() {
	searchCommand.PersistentFlags().BoolVarP(
		&options.search.params, "params", "P", false,
		"list valid search params")
	searchCommand.PersistentFlags().StringVar(
		&options.search.imdbID, "imdbId", "", "search by IMDB ID")
	searchCommand.PersistentFlags().StringVar(
		&options.search.zap2itID, "zap2itId", "", "search by Zap2It ID")
	rootCommand.PersistentFlags().BoolVarP(
		&options.json, "json", "J", false, "output JSON")
	rootCommand.PersistentFlags().StringVarP(
		&options.template, "template", "T", "", "use a custom output template")
	rootCommand.PersistentFlags().BoolVar(
		&options.trace, "trace", false, "trace execution")
	rootCommand.AddCommand(languagesCommand, searchCommand, seriesCommand)

	cmd.AddCommand(rootCommand)
}

// command is a cobra.Command which is to be added to the root command.
var rootCommand = &cobra.Command{
	Use:   "tvdb",
	Short: "interact with the TVDB",
	Long: `
A basic client for working with The TVDB (https://thetvdb.com).
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Usage(); err != nil {
			log.Fatal(err)
		}
	},
}

var languagesCommand = &cobra.Command{
	Use:   "languages",
	Short: "list supported languages",
	Long:  "list all languages supported by The TVDB",
	Run: func(cmd *cobra.Command, args []string) {
		languages, err := DefaultClient.Languages()
		if err != nil {
			log.Fatal(err)
		}
		output(cmd, "languages", languages)
	},
}

// searchCommand searches for a series.
var searchCommand = &cobra.Command{
	Use:   "search TERM",
	Short: "search for things",
	Long:  "search for series on The TVDB",
	Run: func(cmd *cobra.Command, args []string) {
		if options.search.params {
			if len(args) > 0 {
				usage(cmd)
				return
			}
			result, err := DefaultClient.SearchSeriesParams()
			if err != nil {
				log.Fatal(err)
			}
			output(cmd, "search", result)
			return
		}

		values := &url.Values{}
		if options.search.imdbID == "" {
			if options.search.zap2itID == "" {
				values.Set("name", strings.Join(args, " "))
			} else {
				if len(args) != 1 {
					usage(cmd)
					return
				}
				values.Set("zap2itId", args[0])
			}
		} else {
			if len(args) != 1 {
				usage(cmd)
				return
			}
			values.Set("imdbId", args[0])
		}
		result, err := DefaultClient.SearchSeries(values)
		if err != nil {
			log.Fatal(err)
			return
		}
		output(cmd, "search", result)

	},
}

// seriesCommand searches for a series.
var seriesCommand = &cobra.Command{
	Use:   "series ID",
	Short: "get series information",
	Long:  "Gets detailed information for the series with the given TVDB ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			usage(cmd)
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
			return
		}
		result, err := DefaultClient.GetSeries(id)
		if err != nil {
			log.Fatal(err)
			return
		}
		output(cmd, "series", result)
	},
}

var templateFuncs = template.FuncMap{
	"headers": func() (bool, error) { return true, nil },
}

func output(cmd *cobra.Command, t string, data interface{}) {
	if options.json {
		outputJSON(data)
	}
	text := options.template
	if options.template == "" {
		text = templates[t]
	}
	templ, err := template.New(t).Funcs(templateFuncs).Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	if err := templ.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

func outputJSON(data interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", jsonIndentation)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		log.Fatal(err)
	}
}

func usage(cmd *cobra.Command) {
	if err := cmd.Usage(); err != nil {
		log.Fatal(err)
	}
}
