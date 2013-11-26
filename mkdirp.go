package mkdirp

import (
	"regexp"
	"strings"
)

var validPath = regexp.MustCompile(`^.+/$`)

type dirTree struct {
	name string
	subs []*dirTree
}

func mk(path string) []string {
	trimmed := strings.Trim(path, "/ ")
	return []string{path}
}
