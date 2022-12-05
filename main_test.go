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
	"testing"

	"github.com/stretchr/testify/assert"
)

// sort.Interface for []shakableFile based on the filepath field.
type ByPath []shakableFile

func (p ByPath) Len() int           { return len(p) }
func (p ByPath) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPath) Less(i, j int) bool { return p[i].filepath < p[j].filepath }

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
	paths := collect(filesToCollect, []string{}, false)

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
	paths := collect(filesToCollect, []string{filepath.Base(file2.Name())}, false)

	assert.Equal(
		t,
		paths,
		[]shakableFile{
			shakableFile{filepath: file1.Name(), isShaked: false},
			shakableFile{filepath: file3.Name(), isShaked: false}})
}

func TestCollectDirectory(t *testing.T) {

	file1, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file1.Name())
	file2, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file2.Name())
	dir, _ := ioutil.TempDir("", "mvshaker")
	defer os.Remove(dir)

	filesToCollect := []string{file1.Name(), file2.Name(), dir}
	paths := collect(filesToCollect, []string{}, false)

	assert.Equal(
		t,
		paths,
		[]shakableFile{
			shakableFile{filepath: file1.Name(), isShaked: false},
			shakableFile{filepath: file2.Name(), isShaked: false}})
}

func TestCollectDirectoryRecursive(t *testing.T) {

	file1, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file1.Name())
	file2, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file2.Name())
	dir, _ := ioutil.TempDir("", "mvshaker")
	defer os.Remove(dir)
	file3, _ := ioutil.TempFile(dir, "mvshaker")
	defer os.Remove(file3.Name())
	dir2, _ := ioutil.TempDir(dir, "mvshaker")
	defer os.Remove(dir2)
	file4, _ := ioutil.TempFile(dir2, "mvshaker")
	defer os.Remove(file4.Name())

	filesToCollect := []string{file1.Name(), file2.Name(), dir}
	paths := collect(filesToCollect, []string{}, true)
	sort.Sort(ByPath(paths))

	expected := []shakableFile{
		shakableFile{filepath: file1.Name(), isShaked: false},
		shakableFile{filepath: file2.Name(), isShaked: false},
		shakableFile{filepath: file3.Name(), isShaked: false},
		shakableFile{filepath: file4.Name(), isShaked: false}}
	sort.Sort(ByPath(expected))

	assert.Equal(t, paths, expected)
}

func TestShake(t *testing.T) {

	file1, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file1.Name())
	file2, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file2.Name())
	file3, _ := ioutil.TempFile("", "mvshaker")
	defer os.Remove(file3.Name())

	if err := ioutil.WriteFile(file1.Name(), []byte("01\n"), 0644); err != nil {
		assert.Error(t, err)
	}
	if err := ioutil.WriteFile(file2.Name(), []byte("02\n"), 0644); err != nil {
		assert.Error(t, err)
	}
	if err := ioutil.WriteFile(file3.Name(), []byte("03\n"), 0644); err != nil {
		assert.Error(t, err)
	}

	source := []shakableFile{
		shakableFile{filepath: file1.Name(), isShaked: false},
		shakableFile{filepath: file2.Name(), isShaked: false},
		shakableFile{filepath: file3.Name(), isShaked: false}}

	dest := make([]shakableFile, len(source))
	copy(dest, source)

	rand.Seed(1)
	ランダム(dest)

	shake(source, dest)

	assert.Equal(t, source[0].isShaked, false)
	assert.Equal(t, source[1].isShaked, true)
	assert.Equal(t, source[2].isShaked, true)

	assert.Equal(t, dest[0].isShaked, false)
	assert.Equal(t, dest[1].isShaked, true)
	assert.Equal(t, dest[2].isShaked, true)

	data1, err := ioutil.ReadFile(file1.Name())
	if err != nil {
		assert.Error(t, err)
	}
	data2, err := ioutil.ReadFile(file2.Name())
	if err != nil {
		assert.Error(t, err)
	}
	data3, err := ioutil.ReadFile(file3.Name())
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, string(data1), "03\n")
	assert.Equal(t, string(data2), "01\n")
	assert.Equal(t, string(data3), "02\n")
}
