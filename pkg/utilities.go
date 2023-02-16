package pkg

import (
	"crypto/rand"
	"errors"
	rand2 "math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CreateOrOpenFile(filename string) (*os.File, error) {
	if FileExists(filename) {
		return os.OpenFile(filename,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	return os.Create(filename)
}

func FileExists(fn string) bool {
	if _, err := os.Stat(fn); err == nil {
		return true
	}
	return false
}

func Slugify(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string is not permitted")
	}

	var re = regexp.MustCompile(`[^a-z\d]+`)
	slug := strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
	if len(slug) == 0 {
		return "", errors.New("after removing characters, slug iis zero length")
	}
	return slug, nil
}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string
func RandomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}

func RandomNumber(length int) string {
	rand2.Seed(time.Now().Unix())
	var retStr string
	for i := 0; i < length; i++ {
		retStr += strconv.Itoa(rand2.Intn(9))
	}
	return retStr
}
