package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var punctuations = []rune(`!#$%&'()*+,-./:;<=>?@[\]^_{|}~\` + "`")
var kwdBlackList = []string{"and", "else", "you", "from", "like", "where", "delete", "update", "insert", "other", "for"}

func Capitalize(s *string) {
	*s = cases.Title(language.Und).String(*s)
}

func Lower(s *string) {
	*s = strings.ToLower(*s)
}

func Upper(s *string) {
	*s = strings.ToUpper(*s)
}

func Strip(s *string) {
	*s = strings.TrimSpace(*s)
}

func GenKWDs(s string, withHash bool) []string {
	kwds := []string{}
	for _, p := range punctuations {
		s = strings.ReplaceAll(s, string(p), " ")
	}
	s = strings.ToLower(s)
	ss := strings.Split(s, " ")
	for _, _s := range ss {
		if len([]rune(_s)) >= 3 {
			inBlk := false
			for _, __s := range kwdBlackList {
				if _s == __s {
					inBlk = true
					break
				}
			}
			if !inBlk {
				if withHash {
					_s = "#" + _s
				}
				kwds = append(kwds, _s)
			}
		}
	}
	return kwds
}

func CleanTags(t *[]string) []string {
	tags := []string{}
	if t == nil {
		return tags
	}
	for _, _s := range *t {
		Strip(&_s)
		Lower(&_s)
		for _, p := range punctuations {
			_s = strings.ReplaceAll(_s, string(p), "")
		}
		if len([]rune(_s)) >= 3 {
			inBlk := false
			for _, __s := range kwdBlackList {
				if _s == __s {
					inBlk = true
					break
				}
			}
			if !inBlk {
				tags = append(tags, _s)
			}
		}
	}
	return tags
}
