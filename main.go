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

type shakableFile struct {
	filepath string
	isShaked bool
}

func main() {

	var (
		sources = kingpin.Arg("source",
			"Files or directories to skake.").Required().Strings()

		paths []shakableFile
	)

	kingpin.Version(version)
	kingpin.CommandLine.Help = "File shaker for the Masses."
	kingpin.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	for _, source := range *sources {

		file, err := filepath.Abs(source)
		if err != nil {
			continue
		}
		paths = append(paths, shakableFile{filepath: file, isShaked: false})
	}

	dest := make([]shakableFile, len(paths))
	copy(dest, paths)

	ランダム(dest)

	shake(paths, dest)

}

func ランダム(slice []shakableFile) {

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
