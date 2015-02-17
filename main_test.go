package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_E2E_Urls(t *testing.T) {
	assert := assert.New(t)

	a := "radiohead"
	c := "urls"

	w := &bytes.Buffer{}

	Run(a, c, w)

	res := w.String()

	assert.Contains(res, "You have selected the artist radiohead and the category urls.")

}

func Test_E2E_Location(t *testing.T) {
	assert := assert.New(t)

	a := "radiohead"
	c := "location"

	w := &bytes.Buffer{}

	Run(a, c, w)

	res := w.String()

	assert.Contains(res, "You have selected the artist radiohead and the category location.")

}
