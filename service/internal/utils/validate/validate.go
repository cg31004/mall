package validate

import (
	"regexp"
)

func ValidateNumLetter(validateStr string) bool {
	m, _ := regexp.MatchString("^[a-zA-Z0-9]+$", validateStr)
	return m
}
