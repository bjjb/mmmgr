package tvdb

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/bjjb/mmmgr/cmd"
	"github.com/spf13/cobra"
)

var jsonIndentation = "  "

func init() {
	searchCommand.PersistentFlags().BoolP(
		"params", "P", false, "list valid search params")
	searchCommand.PersistentFlags().StringP(
		"name", "n", "", "search by name (default no flags given)")
	searchCommand.PersistentFlags().String(
		"imdbId", "", "search by IMDB ID")
	searchCommand.PersistentFlags().String(
		"zap2itId", "", "search by Zap2It ID")
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

// languagesCommand GETs /languages.
var languagesCommand = &cobra.Command{
	Use:   "languages",
	Short: "list supported languages",
	Long:  "list all languages supported by The TVDB",
	Run: func(cmd *cobra.Command, args []string) {
		languages, err := DefaultClient.Languages()
		if err != nil {
			log.Fatal(err)
		}
		outputJSON(languages)
	},
}

// searchCommand searches for a series.
var searchCommand = &cobra.Command{
	Use:   "search",
	Short: "search for things",
	Long:  "search for series on The TVDB",
	Run: func(cmd *cobra.Command, args []string) {
		params := getBool(cmd, "params")

		if params {
			if len(args) > 0 {
				usage(cmd)
				return
			}
			result, err := DefaultClient.SearchSeriesParams()
			if err != nil {
				log.Fatal(err)
			}
			outputJSON(result)
			return
		}

		values := getValues(cmd, "name", "imdbId", "zap2itId")
		if strings.Contains(values.Encode(), "&") {
			log.Fatal("--name, --imdbId and --zap2itId are mutually exclusive")
			return
		} else if values.Encode() == "" {
			if len(args) == 0 {
				usage(cmd)
				return
			}
			values.Set("name", strings.Join(args, " "))
		} else {
			if len(args) > 0 {
				usage(cmd)
				return
			}
		}

		result, err := DefaultClient.SearchSeries(values)
		if err != nil {
			log.Fatal(err)
			return
		}
		outputJSON(result)
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
		outputJSON(result)
	},
}

// outputJSON simply encodes the data to JSON and prints it to STDOUT
func outputJSON(data interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", jsonIndentation)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		log.Fatal(err)
	}
}

// usage is a utility method for printing a command and panicking if it fails
func usage(cmd *cobra.Command) {
	if err := cmd.Usage(); err != nil {
		log.Fatal(err)
	}
}

// getString is a utility method for getting a string from a command's
// persistent flags, and panicking if it fails.
func getString(cmd *cobra.Command, f string) string {
	s, err := cmd.PersistentFlags().GetString(f)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// getBool is a utility method for getting a boolean from a command's
// persistent flags, and panicking if it fails.
func getBool(cmd *cobra.Command, f string) bool {
	b, err := cmd.PersistentFlags().GetBool(f)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

// getValues is a utility method to extract a series of String flags from the
// commands's persistent flags, and (if they're not empty) add them to a
// url.Values.
func getValues(cmd *cobra.Command, flags ...string) *url.Values {
	values := &url.Values{}
	for _, k := range flags {
		v := getString(cmd, k)
		if v != "" {
			values.Set(k, v)
		}
	}
	return values
}
