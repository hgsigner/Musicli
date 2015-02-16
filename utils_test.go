package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Arg_Not_A_Slice(t *testing.T) {
	assert := assert.New(t)

	slice := 1

	contains := SliceContains(slice, "songs")

	assert.False(contains)

}

func Test_Slice_Contains_String(t *testing.T) {
	assert := assert.New(t)

	slice := []string{
		"urls",
		"songs",
		"albums",
	}

	contains := SliceContains(slice, "songs")

	assert.True(contains)

}

func Test_Slice_Not_Contains_String(t *testing.T) {
	assert := assert.New(t)

	slice := []string{
		"urls",
		"albums",
	}

	contains := SliceContains(slice, "songs")

	assert.False(contains)

}

func Test_Slice_Contains_Int(t *testing.T) {
	assert := assert.New(t)

	slice := []int{1, 2, 3}

	contains := SliceContains(slice, 2)

	assert.True(contains)

}

func Test_Slice_Not_Contains_Int(t *testing.T) {
	assert := assert.New(t)

	slice := []int{1, 2, 3}

	contains := SliceContains(slice, 5)

	assert.False(contains)

}

func Test_CheckInterfaceType(t *testing.T) {

	assert := assert.New(t)
	_, ok := CheckInterfaceType([]int{1, 3}, reflect.Slice)
	assert.True(ok)

}

func Test_Different_CheckInterfaceType(t *testing.T) {

	assert := assert.New(t)
	_, ok := CheckInterfaceType([]int{1, 3}, reflect.String)
	assert.False(ok)

}
