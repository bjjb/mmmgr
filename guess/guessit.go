package guess

import (
	"encoding/json"
	"github.com/bjjb/mmmgr/guessit"
	"log"
)

func Guessit(paths []string) <-chan *Guess {
	out := make(chan *Guess)
	go func() {
		for _, path := range paths {
			data, err := guessit.Guessit(path)
			if err != nil {
				log.Print(err)
				return
			}
			g := New(path)
			if err := json.Unmarshal(data, g); err != nil {
				log.Print(err)
				return
			}
			out <- g
		}
		close(out)
	}()
	return out
}
