package cmd

import (
	"fmt"
	"github.com/bjjb/mmmgr/file"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [files]",
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
		for _, path := range args {
			file := file.New(path)
			fmt.Printf("%#v\n", file)
		}
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)
}
