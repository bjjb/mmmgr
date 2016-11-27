package video

import (
	"github.com/bjjb/mmmgr/files"
	"github.com/bjjb/mmmgr/guess"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	v := New("foo")
	assertEqual(t, "File", v.File, files.New("foo"))
	assertEqual(t, "Languages", v.Languages, []string{"", "", "", "", ""})
	assertEqual(t, "SubtitleLanguages", v.SubtitleLanguages, make([]string, 10))
}

func TestNewFromGuess(t *testing.T) {
	g := guess.New("foo.avi")
	g.Format = "DVD"
	g.VideoCodec = "XviD"
	g.Container = "avi"
	g.Title = "Home Movie"
	languages := []string{"en", "fr", "it"}
	g.Languages = languages[:]

	v := NewFromGuess(g)

	assertEqual(t, "Format", v.Format, "DVD")
	assertEqual(t, "Codec", v.Codec, "XviD")
	assertEqual(t, "Container", v.Container, "avi")
	assertEqual(t, "Title", v.Title, "Home Movie")
	assertEqual(t, "Languages", v.Languages, languages)
}

func assertEqual(t *testing.T, msg string, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s: expected %q, got %q", msg, b, a)
	}
}
