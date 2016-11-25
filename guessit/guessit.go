package guessit

import (
	"log"
	"os/exec"
)

type json []byte

var cmd []string
var cache map[string]json

func init() {
	cache = make(map[string]json)
	if _, err := exec.LookPath("guessit"); err == nil {
		cmd = []string{"guessit", "-j"}
		return
	}
	if _, err := exec.LookPath("docker"); err == nil {
		cmd = []string{"docker", "run", "--rm", "toilal/guessit", "-j"}
		return
	}
	log.Fatal(`Failed to find guessit or docker on the system.
If you have python and pip (or pip3), install it with 'pip install guessit'.
Otherwise install docker on your system, and I can use a docker image
instead.`)
}

func guessit(path string) (json, error) {
	return exec.Command(cmd[0], append(cmd[1:], path)...).Output()
}

func Guessit(path string) json {
	if j, found := cache[path]; found {
		return j
	}
	j, err := guessit(path)
	if err != nil {
		log.Fatalf("guessit error: %q", err)
	}
	cache[path] = j
	return j
}
