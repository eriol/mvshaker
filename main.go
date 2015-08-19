package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

const version = "0.1"

type shakableFile struct {
	filepath string
	isShaked bool
}

func main() {

	const (
		sourcesHelp      = "Files to skake."
		excludeFilesHelp = `Files to be excluded. Use " for multiple arguments and space as separator, e.g. "bash ls".`
	)

	var (
		sources      = kingpin.Arg("source", sourcesHelp).Required().Strings()
		excludeFiles = kingpin.Flag("exclude", excludeFilesHelp).Short('e').String()
	)

	kingpin.Version(version)
	kingpin.CommandLine.Help = "File shaker for the Masses."
	kingpin.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	exclude := strings.Split(*excludeFiles, " ")

	paths := collect(*sources, exclude)

	dest := make([]shakableFile, len(paths))
	copy(dest, paths)

	ランダム(dest)

	shake(paths, dest)

}

func collect(sources, exclude []string) []shakableFile {

	var paths []shakableFile
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
			continue
		}

		target := filepath.Base(path)
		i := sort.SearchStrings(exclude, target)
		if i < len(exclude) && exclude[i] == target {
			continue
		}

		paths = append(paths, shakableFile{filepath: path, isShaked: false})
	}

	return paths

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