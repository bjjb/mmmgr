package guess

import "testing"

func TestFromPath(t *testing.T) {
	cases := []struct{ in, title, t string }{
		{"foo.bar", "foo bar", "movie"},
		{"Witch, The", "The Witch", "movie"},
	}
	for _, c := range cases {
		g := FromPath(c.in)
		if g.Title != c.title {
			t.Errorf("FromPath(%q).Title => %q != %q", c.in, g.Title, c.title)
		}
		if g.Type != c.t {
			t.Errorf("FromPath(%q).Type => %q != %q", c.in, g.Type, c.t)
		}
	}
}
