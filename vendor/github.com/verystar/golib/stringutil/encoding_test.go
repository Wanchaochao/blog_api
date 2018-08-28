package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	originstr = "abcd您好!"
	base64str = "YWJjZOaCqOWlvSE="
)

func TestBase64Encode(t *testing.T) {
	s := Base64Encode([]byte(originstr))
	assert.Equal(t, s, base64str, "Base64Encode must be %s but return %s", base64str, s)
}

func TestBase64Decode(t *testing.T) {
	s, err := Base64Decode(base64str)

	assert.NoError(t, err, "Base64Decode error %s", err)
	assert.Equal(t, string(s), originstr, "Base64Decode must be %s but return %s", originstr, string(s))
}
