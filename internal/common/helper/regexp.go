package helper

import "regexp"

var (
	Regexp      regexper
	regExpEmail = regexp.MustCompile(`^(\.*[a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+`)
)

type regexper struct{}

func (regexper) IsEmail(email string) bool {
	return regExpEmail.MatchString(email)
}
