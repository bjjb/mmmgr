package guess

import (
	"encoding/json"
	"github.com/bjjb/mmmgr/guessit"
	"log"
)

type Guess struct {
	AlternativeTitle []string `json:"alternative_title"`
	AudioChannels    string   `json:"audio_channels"`
	AudioCodec       string   `json:"audio_codec"`
	AudioProfile     string   `json:"audio_profile"`
	Bonus            string   `json:"bonus"`
	BonusTitle       string   `json:"bonus_title"`
	CD               string   `json:"cd"`
	CDCount          string   `json:"cd_count"`
	Container        string   `json:"container"`
	Country          string   `json:"country"`
	CRC32            string   `json:"crc32"`
	Date             string   `json:"date"`
	Edition          string   `json:"edition"`
	Episode          int      `json:"episode"`
	EpisodeTitle     string   `json:"episode_title"`
	EpisodeCount     int      `json:"episode_count"`
	EpisodeDetails   string   `json:"episode_details"`
	EpisodeFormat    string   `json:"episode_format"`
	Film             string   `json:"film"`
	FilmTitle        string   `json:"film_title"`
	Format           string   `json:"format"`
	Language         string   `json:"language"`
	MimeType         string   `json:"mime_type"`
	Other            string   `json:"other"`
	Part             int      `json:"part"`
	ProperCount      string   `json:"proper_count"`
	ReleaseGroup     string   `json:"release_group"`
	ScreenSize       string   `json:"screen_size"`
	Season           int      `json:"season"`
	SeasonCount      int      `json:"season_count"`
	SubtitleLanguage string   `json:"subtitle_language"`
	Title            string   `json:"title"`
	Type             string   `json:"type"`
	UUID             string   `json:"uuid"`
	Version          string   `json:"version"`
	VideoAPI         string   `json:"video_api"`
	VideoCodec       string   `json:"video_codec"`
	VideoProfile     string   `json:"video_profile"`
	Website          string   `json:"website"`
	Year             int      `json:"year"`
}

func New(path string) *Guess {
	g := new(Guess)
	data := guessit.Guessit(path)
	err := json.Unmarshal(data, &g)
	if err != nil {
		log.Fatalf("error decoding guessit result: %q %v", err, data)
	}
	return g
}
