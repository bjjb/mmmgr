package file

import "testing"

func TestIsMovie(t *testing.T) {
	cases := []struct {
		t string
		r bool
	}{
		{"Rambo.avi", true},
		{"Rambo.1991.avi", true},
	}
	for _, c := range cases {
		if r := IsMovie(c.t); r != c.r {
			t.Errorf("IsMovie(%q) => %v, expected %v", c.t, r, c.r)
		}
	}
}

func TestIsTV(t *testing.T) {
	cases := []struct {
		t string
		r bool
	}{
		{"Mr.Robot.S01E11.BlahBlah.avi", true},
		{"Rambo.1991.avi", false},
	}
	for _, c := range cases {
		if r := IsTV(c.t); r != c.r {
			t.Errorf("IsTV(%q) => %v, expected %v", c.t, r, c.r)
		}
	}
}
