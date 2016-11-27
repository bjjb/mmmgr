package guessit

import (
	"github.com/bjjb/mmmgr/assert"
	"strings"
	"testing"
)

func Test_guessit(t *testing.T) {
	squeezer := strings.NewReplacer(" ", "", "\n", "")
	testCases := []struct{ in, ok string }{
		{"hello", `{"title":"hello","type":"movie"}`},
		{"hello.1999", `{"title":"hello","year":1999,"type":"movie"}`},
		{"hello.world.1999[yarbag].HDTV.XviD.mkv", `{"title":"helloworld","year":1999,"release_group":"yarbag","format":"HDTV","video_codec":"XviD","container":"mkv","mimetype":"video/x-matroska","type":"movie"}`},
		{"s01e01.hello.1999", `{"season":1,"episode":1,"title":"hello","year":1999,"type":"episode"}`},
	}
	for _, c := range testCases {
		result, err := guessit(c.in)
		if err != nil {
			t.Error(err)
		}
		out := squeezer.Replace(string(result))
		assert.Equal(t, c.ok, out)
	}
}
