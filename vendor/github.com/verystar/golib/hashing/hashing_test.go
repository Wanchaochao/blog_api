package hashing

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	str   = "abcd"
	signs = map[string]string{
		"md5":     "e2fc714c4727ee9395f324cd2e7f331f",
		"sha1":    "81fe8bfe87576c3ecb22426f8e57847382917acf",
		"sha256":  "88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589",
		"sha512":  "d8022f2060ad6efd297ab73dcc5355c9b214054b0d1776a136a669d26a7d3b14f73aa0d0ebff19ee333368f0164b6419a96da49e3e481753e7e96b716bdccb6f",
		"signmd5": "1ba488f6003e87178d16c1ae47f4aee9",
	}
)

func TestMd5(t *testing.T) {
	sign := Md5(str)
	assert.Equal(t, sign, signs["md5"], "md5 hasing must be %s but return %s", sign, signs["md5"])
}

func TestSha1(t *testing.T) {
	sign := Sha1(str)
	assert.Equal(t, sign, signs["sha1"], "sha1 hasing must be %s but return %s", sign, signs["sha1"])
}

func TestSha256(t *testing.T) {
	sign := Sha256(str)
	assert.Equal(t, sign, signs["sha256"], "sha256 hasing must be %s but return %s", sign, signs["sha256"])
}

func TestSha512(t *testing.T) {
	sign := Sha512(str)
	assert.Equal(t, sign, signs["sha512"], "sha512 hasing must be %s but return %s", sign, signs["sha512"])
}

func TestSign(t *testing.T) {
	values := url.Values{}
	values.Add("id", "1")
	values.Add("name", "您好")
	values.Add("date", "2012-10-24")
	values.Add("status", "0")
	values.Add("sign", "1fdslfi2o34nldsfl")

	sign := Sign(values, "abcd")
	assert.Equal(t, sign, signs["signmd5"], "signmd5 hasing must be %s but return %s", sign, signs["signmd5"])

}
