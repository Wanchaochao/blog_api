package stringutil

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

//生成随机字符串
const randomStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//5倍的性能提升，放到外面
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = randomStr[seededRand.Intn(len(randomStr))]
	}
	return string(b)
}

func TitleCasedName(name string) string {
	newstr := make([]rune, 0)
	upNextChar := true

	name = strings.ToLower(name)

	for _, chr := range name {
		switch {
		case upNextChar:
			upNextChar = false
			if 'a' <= chr && chr <= 'z' {
				chr -= 'a' - 'A'
			}
		case chr == '_':
			upNextChar = true
			continue
		}

		newstr = append(newstr, chr)
	}

	return string(newstr)
}

func SnakeCasedName(name string) string {
	newstr := make([]rune, 0)
	for idx, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '_')
			}
			chr -= 'A' - 'a'
		}
		newstr = append(newstr, chr)
	}

	return string(newstr)
}

func TrimBom(s string) string {
	buf := []byte(s)
	if len(buf) > 3 {
		//0xef, 0xbb, 0xbf 239 187 191
		if buf[0] == 239 && buf[1] == 187 && buf[2] == 191 {
			return string(buf[3:])
		}
	}
	return s
}

// Leftpad
func Leftpad(s string, length int, ch ...rune) string {
	c := ' '
	if len(ch) > 0 {
		c = ch[0]
	}
	l := length - utf8.RuneCountInString(s)
	if l > 0 {
		s = strings.Repeat(string(c), l) + s
	}
	return s
}