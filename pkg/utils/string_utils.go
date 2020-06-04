package utils

import (
	"regexp"
	"strings"
)

// ReplaceSpecialChar replaces any non-alphanumeric character in a string with a character of choice
func ReplaceSpecialChar(str string, rep string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9\\s]+")
	return reg.ReplaceAllString(str, rep)
}

// ToSnakeCase turns camel/pascal case strings into snake_case strings
func ToSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	matchAllSpaces := regexp.MustCompile("(\\s)")
	cleanUpHack := regexp.MustCompile("i_ds")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = matchAllSpaces.ReplaceAllString(snake, "_")
	snake = strings.ToLower(snake)
	snake = cleanUpHack.ReplaceAllString(snake, "ids")

	return snake
}
