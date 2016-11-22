package music

import "testing"

func TestGuessFromPath(t *testing.T) {
	cases := []struct{ in, title string }{
		{"blah.mp3", "blah"},
	}
	for _, c := range cases {
		if r := GuessFromPath(c.in).Title; r != c.title {
			t.Errorf("GuessFromPath(%q), expected %q, got %q", c.in, c.title, r)
		}
	}
}
