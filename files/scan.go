package files

import (
	"github.com/bjjb/mmmgr/guess"
	"os"
	"path/filepath"
)

// Scans the given directory for media files
func Scan(path string) <-chan *File {
	out := make(chan *File)
	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if guess.Type(path) == "" {
			return nil
		}
		out <- New(path)
		return nil
	}
	go func() {
		filepath.Walk(path, walker)
		close(out)
	}()
	return out
}
