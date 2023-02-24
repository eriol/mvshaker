// Copyright © 2015 Daniele Tricoli <eriol@mornie.org>.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main // import "noa.mornie.org/eriol/mvshaker"

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/alecthomas/kingpin/v2"
)

const version = "0.3"

type shakableFile struct {
	filepath string
	isShaked bool
}

func main() {

	const (
		sourcesHelp      = "Files to skake."
		excludeFilesHelp = "Files to be excluded. Call it multiple time to exclude more than one file, e.g. -e bash -e ls."
		recursiveHelp    = "Include recursively files inside directories."
	)

	var (
		sources   = kingpin.Arg("source", sourcesHelp).Required().Strings()
		exclude   = kingpin.Flag("exclude", excludeFilesHelp).Short('e').Strings()
		recursive = kingpin.Flag("recursive", recursiveHelp).Short('r').Bool()
	)

	kingpin.Version(version)
	kingpin.CommandLine.Help = "File shaker for the Masses."
	kingpin.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	paths := collect(*sources, *exclude, *recursive)

	dest := make([]shakableFile, len(paths))
	copy(dest, paths)

	ランダム(dest)

	shake(paths, dest)

}

func collect(sources, exclude []string, recursive bool) []shakableFile {

	var paths []string
	var files []shakableFile

	sort.Strings(exclude)

	for _, source := range sources {

		path, err := filepath.Abs(source)
		if err != nil {
			continue
		}

		fi, err := os.Stat(path)
		if err != nil {
			continue
		}

		if fi.IsDir() {
			if recursive {
				filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
					if !info.IsDir() {
						paths = append(paths, path)
					}
					return nil
				})
				continue
			} else {
				continue
			}
		}

		paths = append(paths, path)
	}

	for _, path := range paths {
		target := filepath.Base(path)
		i := sort.SearchStrings(exclude, target)
		if i < len(exclude) && exclude[i] == target {
			continue
		}
		files = append(files, shakableFile{filepath: path, isShaked: false})
	}
	return files

}

func ランダム(slice []shakableFile) {

	// Reverse traversing because rand.Intn panics if argument is <= 0.
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i)
		slice[i], slice[j] = slice[j], slice[i]
	}

}

func shake(source, dest []shakableFile) {

	for i := len(source) - 1; i > 0; i-- {
		if source[i].isShaked == false || dest[i].isShaked == false {

			file, err := ioutil.TempFile(filepath.Dir(source[i].filepath),
				"mvshaker")
			defer os.Remove(file.Name())

			if err != nil {
				panic(err)
			}

			if err := os.Rename(source[i].filepath, file.Name()); err != nil {
				panic(err)
			}
			if err := os.Rename(dest[i].filepath, source[i].filepath); err != nil {
				panic(err)
			}
			if err := os.Rename(file.Name(), dest[i].filepath); err != nil {
				panic(err)
			}

			source[i].isShaked = true
			dest[i].isShaked = true
		}
	}
}
