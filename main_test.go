package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testランダム(t *testing.T) {
	files := []shakableFile{
		shakableFile{filepath: "a.txt", isShaked: false},
		shakableFile{filepath: "b.txt", isShaked: false},
		shakableFile{filepath: "c.txt", isShaked: false}}

	assert.Equal(t, files[0].filepath, "a.txt")
	assert.Equal(t, files[1].filepath, "b.txt")
	assert.Equal(t, files[2].filepath, "c.txt")

	rand.Seed(1)

	ランダム(files)

	assert.Equal(t, files[0].filepath, "c.txt")
	assert.Equal(t, files[1].filepath, "a.txt")
	assert.Equal(t, files[2].filepath, "b.txt")

}

func TestCollectNoExclude(t *testing.T) {

	file1, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file1.Name())
	file2, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file2.Name())
	file3, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file3.Name())

	filesToCollect := []string{file1.Name(), file2.Name(), file3.Name()}
	paths := collect(filesToCollect, []string{})

	assert.Equal(
		t,
		paths,
		[]shakableFile{
			shakableFile{filepath: file1.Name(), isShaked: false},
			shakableFile{filepath: file2.Name(), isShaked: false},
			shakableFile{filepath: file3.Name(), isShaked: false}})
}

func TestCollectExclude(t *testing.T) {

	file1, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file1.Name())
	file2, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file2.Name())
	file3, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file3.Name())

	filesToCollect := []string{file1.Name(), file2.Name(), file3.Name()}
	paths := collect(filesToCollect, []string{filepath.Base(file2.Name())})

	assert.Equal(
		t,
		paths,
		[]shakableFile{
			shakableFile{filepath: file1.Name(), isShaked: false},
			shakableFile{filepath: file3.Name(), isShaked: false}})
}
