package hashing

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"net/url"
	"sort"
	"strings"
)

func encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func Sign(values url.Values, app_secret string, hash ...string) (string) {
	hashName := "md5"
	if len(hash) > 0 {
		hashName = strings.ToLower(hash[0])
	}

	for key, _ := range values {
		if values.Get(key) == "" || key == "sign" {
			values.Del(key)
		}
	}

	str := encode(values) + app_secret

	switch hashName {
	case "sha1":
		return Sha1(str)
	case "sha256":
		return Sha256(str)
	case "sha512":
		return Sha512(str)
	default:
		return Md5(str)
	}
}

func Md5(text string) string {
	algorithm := md5.New()
	return stringHasher(algorithm, text)
}

// Sha1 hashes using sha1 algorithm
func Sha1(text string) string {
	algorithm := sha1.New()
	return stringHasher(algorithm, text)
}

// Sha256 hashes using sha256 algorithm
func Sha256(text string) string {
	algorithm := sha256.New()
	return stringHasher(algorithm, text)
}

// Sha512 hashes using sha512 algorithm
func Sha512(text string) string {
	algorithm := sha512.New()
	return stringHasher(algorithm, text)
}

func stringHasher(algorithm hash.Hash, text string) string {
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
