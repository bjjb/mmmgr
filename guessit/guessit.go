package guessit

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Guess struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	Show         string `json:"show"`
	EpisodeTitle string `json:"episode_title"`
	Format       string `json:"format"`
	ReleaseGroup string `json:"release_group"`
	Container    string `json:"container"`
	MimeType     string `json:"mimetype"`
	VideoCodec   string `json:"video_codec"`
	Year         int    `json:"year"`
	Season       int    `json:"season"`
	Episode      int    `json:"episode"`
}

// The location of the guessit binary
var exe = "guessit"

func init() {
	if v := os.Getenv("GUESSIT"); v != "" {
		exe = v
	}
}

// Runs guessit -j on the given path, and returns a Guess
func Guessit(path string) *Guess {
	guess := new(Guess)
	out, err := exec.Command(exe, "-j", path).Output()
	if err != nil {
		fmt.Errorf("error from %q: %s", exe, err)
	}
	err = json.Unmarshal(out, &guess)
	return guess
}
