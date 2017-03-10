package tvdb

import (
	"log"
)

/*
An Episode encapsulates a single TVDB TV show episode.
*/
type Episode struct {
	ID                 int      `json:"id"`
	AiredSeason        int      `json:"airedSeason"`
	AiredEpisodeNumber int      `json:"airedEpisodeNumber"`
	EpisodeName        string   `json:"episodeName"`
	FirstAired         string   `json:"firstAired"`
	GuestStars         []string `json:"guestStars"`
	Director           string   `json:"director"`
	Directors          []string `json:"directors"`
	Writers            []string `json:"writers"`
	Overview           string   `json:"overview"`
	ProductionCode     string   `json:"productionCode"`
	ShowURL            string   `json:"showUrl"`
	LastUpdated        int      `json:"lastUpdated"`
	DVDDiscID          string   `json:"dvdDiscid"`
	DVDSeason          int      `json:"dvdSeason"`
	DVDEpisodeNumber   int      `json:"dvdEpisodeNumber"`
	DVDChapter         int      `json:"dvdChapter"`
	AbsoluteNumber     int      `json:"absoluteNumber"`
	Filename           string   `json:"filename"`
	SeriesID           string   `json:"seriesId"`
	LastUpdatedBy      string   `json:"lastUpdatedBy"`
	AirsAfterSeason    int      `json:"airsAfterSeason"`
	AirsBeforeSeason   int      `json:"airsBeforeSeason"`
	AirsBeforeEpisode  int      `json:"airsBeforeEpisode"`
	ThumbAuthor        int      `json:"thumbAuthor"`
	ThumbAdded         string   `json:"thumbAdded"`
	ThumbWidth         string   `json:"thumbWidth"`
	ThumbHeight        string   `json:"thumbHeight"`
	IMDBID             string   `json:"imdbId"`
	SiteRating         int      `json:"siteRating"`
	SiteRatingCount    int      `json:"siteRatingCount"`
}

/*
GetEpisodes gets the Episodes of this Series.
*/
func (s *Series) GetEpisodes() {
	if s.Client == nil {
		log.Fatal("series has no client")
	}
}
