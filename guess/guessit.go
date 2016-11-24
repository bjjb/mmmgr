package guess

import (
	"log"
	"os/exec"
)

func initGuessit() {
	if isExecutable("guessit") {
		guessitCommand = "guessit -j"
		return
	}
	if isExecutable("docker") {
		guessitCommand = "docker run --rm toilal/guessit -j"
		return
	}
	log.Fatal(`Failed to find guessit or docker on the system.
If you have python and pip (or pip3), install it with 'pip install guessit'.
Otherwise install docker on your system, and I can use a docker image
instead.`)
}

func isExecutable(program string) bool {
	_, err := exec.LookPath(program)
	return err == nil
}
