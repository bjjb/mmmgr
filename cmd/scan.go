// Copyright © 2016 JJ Buckley <jj@bjjb.org>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/bjjb/mmmgr/file"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan path",
	Short: "Scan the given path for multimedia",
	Long: `Scan the given path for multimedia files, printing out matches.
For every media file encountered, it simply prints the absolute path. Useful
for chaining commands together in shell scripts. For example:

$ for f in $(mmmgr scan .); do mmmgr add $f; done

The snippet above will add each file to the mmmgr library.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, path := range args {
			filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if err == nil && !info.IsDir() && file.IsMedia(path) {
					fmt.Println(path)
				}
				return err
			})
		}
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)
}
