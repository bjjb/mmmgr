package assert

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func WriteTempFiles(t *testing.T, d string, files []string) {
	content := []byte("File! ✓")
	for _, f := range files {
		path := filepath.Join(d, f)
		if err := ioutil.WriteFile(path, content, 0666); err != nil {
			t.Fatal(err)
		}
	}
}
