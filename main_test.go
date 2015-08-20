package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testランダム(t *testing.T) {
	files := []shakableFile{
		shakableFile{filepath: "a.txt", isShaked: false},
		shakableFile{filepath: "b.txt", isShaked: false},
		shakableFile{filepath: "c.txt", isShaked: false}}

	assert.Equal(t, files[0].filepath, "a.txt", "")
	assert.Equal(t, files[1].filepath, "b.txt", "")
	assert.Equal(t, files[2].filepath, "c.txt", "")

	rand.Seed(1)

	ランダム(files)

	assert.Equal(t, files[0].filepath, "c.txt", "")
	assert.Equal(t, files[1].filepath, "a.txt", "")
	assert.Equal(t, files[2].filepath, "b.txt", "")

}
