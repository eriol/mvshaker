package main

import (
	"math/rand"
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

}

func ランダム(slice []shakedFile) {

	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i)
		slice[i], slice[j] = slice[j], slice[i]
	}

}
