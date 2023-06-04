// inspired from https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6

package cases

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToDottedSnakeCase converts dotted.camelCaseIdentifiers to dotted.snake_case_identifiers
func ToDottedSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

// ToDottedUpperSnakeCase converts dotted.camelCaseIdentifiers to DOTTED.SNAKE_CASE_IDENTIFIERS
func ToDottedUpperSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToUpper(snake)
}
