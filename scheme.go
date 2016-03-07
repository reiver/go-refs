package refs


import (
	"errors"
	"strings"
)


// extractScheme extracts the scheme from a reference string.
//
// For example, if the reference string is "http://example.com/robots.txt",
// the the scheme is "http".
//
// I.e., the scheme is everything before the (first) colon.
func extractScheme(r string) (string, error) {

	colonIndex := strings.IndexRune(r, ':')
	if -1 == colonIndex {
//@TODO: Make a better error.
		return "", errors.New("Bad Request: syntax error; missing colon.")
	}
	scheme := r[:colonIndex]

	return scheme, nil
}
