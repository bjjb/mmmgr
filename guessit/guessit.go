/*
guessit provides a wrapper around toilal/guessit, a python library for
guessing video-file names.
*/

package guessit

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Guessit(path string) ([]byte, error) {
	return guessit(path)
}

var cache map[string][]byte
var cmd []string
var function func(string) ([]byte, error)

func init() {
	cache = make(map[string][]byte)
	switch {
	case isCommand("guessit"):
		function = func(path string) ([]byte, error) {
			return execute("guessit -j", path)
		}
	case isCommand("docker"):
		function = func(path string) ([]byte, error) {
			return execute("docker run --rm toilal/guessit -j", path)
		}
	default:
		function = func(path string) ([]byte, error) {
			return []byte(""), errors.New("guessit not initialized")
		}
	}
}

// Runs guessit (or guessit in a docker)
func guessit(path string) ([]byte, error) {
	// return immediately if the result was cached
	if v, found := cache[path]; found {
		return v, nil
	}
	result, err := function(path)
	if err != nil {
		result = []byte(fmt.Sprintf(`{"error":%q})`, err))
		return result, err
	} else {
		cache[path] = result
		return result, nil
	}
}

// Utility function to test whether a command exists on the system
func isCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// Utility function to execute a command and return the output. Splits the cmd
// and uses the first one as the actual command to exec.Command, and the rest
// (including the args) as arguments.
func execute(argc string, argv ...string) ([]byte, error) {
	args := append(strings.Split(argc, " "), argv...)
	argc = args[0]
	argv = args[1:]
	out, err := exec.Command(argc, argv...).Output()
	if err != nil {
		return []byte("{}"), err
	}
	return out, nil
}
