package guess

import (
	"github.com/bjjb/mmmgr/assert"
	"testing"
)

func TestNew(t *testing.T) {
	g := New("foo")
	assert.Equal(t, "foo", g.Path)
}

func TestGuessAll(t *testing.T) {
}
