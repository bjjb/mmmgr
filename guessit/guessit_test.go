package guessit

import (
	"testing"
)

func TestGuessit(t *testing.T) {
	t.Run("guessit", func(t *testing.T) {
		if _, err := guessit("Hello World"); err != nil {
			t.Fatalf("Error from guessit: %v", err)
		}
	})
}
