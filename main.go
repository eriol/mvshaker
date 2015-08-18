package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

const version = "0.1"

type shakedFile struct {
	filepath string
	shaked   bool
}

func main() {

	var (
		sources = kingpin.Arg("source",
			"Files or directories to skake.").Required().Strings()

		paths []shakedFile
	)

	kingpin.Version(version)
	kingpin.CommandLine.Help = "Files shaker for the Masses."
	kingpin.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	for _, source := range *sources {

		file, err := filepath.Abs(source)
		if err != nil {
			continue
		}
		paths = append(paths, shakedFile{filepath: file, shaked: false})
	}

	dest := make([]shakedFile, len(paths))
	copy(dest, paths)

	ランダム(dest)

	shake(paths, dest)

}

func ランダム(slice []shakedFile) {

	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i)
		slice[i], slice[j] = slice[j], slice[i]
	}

}

func shake(source, dest []shakedFile) {

	for i := len(source) - 1; i > 0; i-- {
		if source[i].shaked == false || dest[i].shaked == false {

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

			source[i].shaked = true
			dest[i].shaked = true
		}
	}
}
