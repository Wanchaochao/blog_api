package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteSize(t *testing.T) {
	assert.EqualValues(t, MustParse("1"), 1)
	assert.EqualValues(t, MustParse("1b"), 1)
	assert.EqualValues(t, MustParse("1k"), KB)
	assert.EqualValues(t, MustParse("1m"), MB)
	assert.EqualValues(t, MustParse("1g"), GB)
	assert.EqualValues(t, MustParse("1t"), TB)
	assert.EqualValues(t, MustParse("1p"), PB)

	assert.EqualValues(t, MustParse(" -1"), -1)
	assert.EqualValues(t, MustParse(" -1 b"), -1)
	assert.EqualValues(t, MustParse(" -1 kb "), -1*KB)
	assert.EqualValues(t, MustParse(" -1 mb "), -1*MB)
	assert.EqualValues(t, MustParse(" -1 gb "), -1*GB)
	assert.EqualValues(t, MustParse(" -1 tb "), -1*TB)
	assert.EqualValues(t, MustParse(" -1 pb "), -1*PB)

	assert.EqualValues(t, MustParse(" 1.5"), 1)
	assert.EqualValues(t, MustParse(" 1.5 kb "), 1.5*KB)
	assert.EqualValues(t, MustParse(" 1.5 mb "), 1.5*MB)
	assert.EqualValues(t, MustParse(" 1.5 gb "), 1.5*GB)
	assert.EqualValues(t, MustParse(" 1.5 tb "), 1.5*TB)
	assert.EqualValues(t, MustParse(" 1.5 pb "), 1.5*PB)
}

func TestByteSizeError(t *testing.T) {
	var err error
	_, err = Parse("--1")
	assert.Equal(t, err, ErrBadByteSize)
	_, err = Parse("hello world")
	assert.Equal(t, err, ErrBadByteSize)
	_, err = Parse("123.132.32")
	assert.Equal(t, err, ErrBadByteSize)
}
