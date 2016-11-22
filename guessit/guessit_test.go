package guessit

import "testing"

func TestGuessit(t *testing.T) {
	cases := []struct{ in, title, t string }{
		{"foo.bar", "foo bar", "movie"},
		{"Witch, The", "The Witch", "movie"},
	}
	for _, c := range cases {
		g := Guessit(c.in)
		if g.Title != c.title {
			t.Errorf("Guessit(%q).Title => %q != %q", c.in, g.Title, c.title)
		}
		if g.Type != c.t {
			t.Errorf("Guessit(%q).Type => %q != %q", c.in, g.Type, c.t)
		}
	}
}
