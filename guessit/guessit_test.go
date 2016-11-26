package guessit

import (
	"strings"
	"testing"
)

func Test_guessit(t *testing.T) {
	squeezer := strings.NewReplacer(" ", "", "\n", "")
	tests := []struct{ in, ok string }{
		{"hello", `{"title":"hello","type":"movie"}`},
		{"hello.1999", `{"title":"hello","year":1999,"type":"movie"}`},
		{"hello.world.1999[yarbag].HDTV.XviD.mkv", `{"title":"helloworld","year":1999,"release_group":"yarbag","format":"HDTV","video_codec":"XviD","container":"mkv","mimetype":"video/x-matroska","type":"movie"}`},
		{"s01e01.hello.1999", `{"season":1,"episode":1,"title":"hello","year":1999,"type":"episode"}`},
	}
	for _, c := range tests {
		out := squeezer.Replace(string(guessit(c.in)))
		if out != c.ok {
			t.Errorf("Guessit(%q): expected %q, got %q", c.in, c.ok, out)
		}
	}
}
