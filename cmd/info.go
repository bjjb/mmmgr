package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/bjjb/mmmgr/audio"
	"github.com/bjjb/mmmgr/books"
	"github.com/bjjb/mmmgr/files"
	"github.com/bjjb/mmmgr/movies"
	"github.com/bjjb/mmmgr/music"
	"github.com/bjjb/mmmgr/tv"
	"github.com/bjjb/mmmgr/video"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/template"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info FILE",
	Short: "Prints information about the given files",
	Long: `Prints out information about the media files.
For all media files, the following info is abailable:

	path	the absolute path to the file on disk
	type	video, movie, tv, audio, music or ebook
	mime	the mime-type of the file

'movie' and 'tv' are subtypes of 'video' for which additional information has
been inferred from the filename, and 'music' is similarly a subtype of
'audio'.

For types 'video', and 'audio', there's always:

	duration	the length of the media in seconds

For 'movie', 'tv', 'music' and 'ebook', there's:

	date	the publication date - sometimes just the year
	title	the media title

'tv' types also have

	show	the TV show name
	season 	the season number, if applicable, 0 if not
	number 	the position within a season, 0 if it's out of sequence
	tvdbid	the ID of the TV episode on TheTVDB.com

'movie' and 'tv also have the extra property

	tmdbid	the ID of the TV episode on TheMovieDB.org

'music' types will have

	artist	recording artist
	album	album name
	number	track number
	disc	disc number (0 if inapplicable)
	mbid	MusicBrainz release ID

Musical information is not always identifiable by these fields alone (though
it's usually enough to organise your catalogue), so you can always delve
further by using the 'mbid' to organise using MusicBrainz (provided you have a
MusicBrainz API key).

Similarly, you can obtain additional info for movies and TV shows using the
TheMovieDB and TheTVDB (check out the 'tmdb', 'tvdb' and 'musicbrainz'
commands).

Information is presented by default in prettified JSON. This can be controlled
by passing the --format flag, allowing you to specify a Go-style template
instead, in which the keys above can be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		output := makeOutputFunction()
		if len(args) == 0 {
			fmt.Println("No files specified")
			os.Exit(1)
		}
		for _, path := range args {
			file := files.New(path)
			switch file.MediaType {
			case "tv":
				output(tv.New(file.Path))
			case "movie":
				output(movies.New(file.Path))
			case "video":
				output(video.New(file.Path))
			case "music":
				output(music.New(file.Path))
			case "audio":
				output(audio.New(file.Path))
			case "book":
				output(books.New(file.Path))
			default:
				output(file)
			}
		}
	},
}

var templ string

func makeOutputFunction() func(interface{}) {
	if templ != "" {
		t := template.Must(template.New("info").Parse(templ))
		return func(x interface{}) {
			if err := t.Execute(os.Stdout, x); err != nil {
				log.Panic(err)
			}
		}
	}
	return outputJSON
}

func outputJSON(x interface{}) {
	json, err := json.MarshalIndent(x, "", "\t")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s\n", json)
}

func init() {
	RootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVar(&templ, "template", "", "output template")
}
